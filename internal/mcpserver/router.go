package mcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Args is the input schema shared by every routed tool.
type Args struct {
	Action string         `json:"action" jsonschema:"the subcommand to execute; must be one of the actions listed in this tool's description"`
	Params map[string]any `json:"params,omitempty" jsonschema:"command flags: keys are flag names without the leading --, values are string, number, boolean, or array (joined with commas). Call geelark_help to see the flags each action accepts"`
}

// action maps an MCP action name to a CLI subcommand.
type action struct {
	key  string   // action name exposed to the LLM
	path []string // CLI command path, e.g. ["phone", "automation"]
	sub  string   // CLI subcommand; empty means same as key
	desc string   // one-line description
}

// act defines an action whose key equals the CLI subcommand name.
func act(path []string, key, desc string) action {
	return action{key: key, path: path, desc: desc}
}

// actAs defines an action whose key differs from the CLI subcommand name.
func actAs(path []string, key, sub, desc string) action {
	return action{key: key, path: path, sub: sub, desc: desc}
}

// routedTool is a tool that dispatches to geelark-cli subcommands.
type routedTool struct {
	name    string
	summary string
	actions []action
}

func (t *routedTool) describe() string {
	var b strings.Builder
	b.WriteString(t.summary)
	b.WriteString("\n\nActions (use geelark_help to get the full flag list of any action):\n")
	for _, a := range t.actions {
		fmt.Fprintf(&b, "- %s: %s\n", a.key, a.desc)
	}
	return b.String()
}

func (t *routedTool) actionList() string {
	keys := make([]string, 0, len(t.actions))
	for _, a := range t.actions {
		keys = append(keys, a.key)
	}
	sort.Strings(keys)
	return strings.Join(keys, ", ")
}

func (t *routedTool) find(key string) *action {
	for i := range t.actions {
		if t.actions[i].key == key {
			return &t.actions[i]
		}
	}
	return nil
}

func (s *Server) registerRouted(srv *mcp.Server, t *routedTool) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        t.name,
		Description: t.describe(),
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args Args) (*mcp.CallToolResult, any, error) {
		return s.handleRouted(ctx, t, args)
	})
}

func (s *Server) handleRouted(ctx context.Context, t *routedTool, args Args) (*mcp.CallToolResult, any, error) {
	a := t.find(args.Action)
	if a == nil {
		return errorResult(fmt.Sprintf(
			"unknown action %q for %s. Valid actions: %s",
			args.Action, t.name, t.actionList(),
		)), nil, nil
	}
	sub := a.sub
	if sub == "" {
		sub = a.key
	}
	oc, err := s.run(ctx, a.path, sub, args.Params)
	if err != nil {
		return errorResult(err.Error()), nil, nil
	}
	cmdPath := strings.Join(append(a.path, sub), " ")
	if !oc.OK() {
		// CLI stderr/stdout contains usage errors; attach the reference doc
		// so the model can fix parameter errors in one round.
		return errorResult(s.withRefDoc(cmdPath, oc.FailureText())), nil, nil
	}
	out := oc.Stdout
	// API-level failure: process exited 0 but the JSON envelope reports a
	// non-zero code. Mark IsError so the model does not treat it as success.
	if c, ok := envelopeCode(out); ok && c != 0 {
		return errorResult(s.withRefDoc(cmdPath, out)), nil, nil
	}
	// Append enum hints from the command's reference doc so any LLM can
	// interpret numeric enum fields (e.g. status: 2) without extra lookups.
	if h := s.docs.EnumHints(cmdPath, out); h != "" {
		out += "\n\n[Enum hints] Numeric field values in the response above:\n" + h
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: out}},
	}, nil, nil
}

// withRefDoc appends the command's reference doc when available.
func (s *Server) withRefDoc(cmdPath, msg string) string {
	if doc, err := s.docs.Get(cmdPath); err == nil {
		return msg + "\n\n[Reference doc for this command]\n" + strings.TrimSpace(doc)
	}
	return msg
}

// envelopeCode extracts the top-level "code" field from a JSON API envelope.
// Returns false when the output is not a JSON object.
func envelopeCode(s string) (int, bool) {
	var e struct {
		Code int `json:"code"`
	}
	if json.Unmarshal([]byte(s), &e) != nil {
		return 0, false
	}
	return e.Code, true
}

func errorResult(msg string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: msg}},
		IsError: true,
	}
}

func allTools() []*routedTool {
	var out []*routedTool
	out = append(out, phoneTools()...)
	out = append(out, phoneAutomationTools()...)
	out = append(out, phoneSystemTools()...)
	out = append(out, browserTools()...)
	out = append(out, miscTools()...)
	return out
}
