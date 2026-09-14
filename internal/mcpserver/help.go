package mcpserver

import (
	"context"
	"regexp"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// registerHelp exposes the full --help text (all flags with descriptions) of
// any geelark-cli command path.
func (s *Server) registerHelp(srv *mcp.Server) {
	mcp.AddTool(srv, &mcp.Tool{
		Name: "geelark_help",
		Description: "Get the CLI help text (flag list with descriptions) for any geelark-cli command, " +
			"e.g. \"phone start\", \"phone automation tiktok-login\", \"browser list\", \"proxy add\". " +
			"Call this before invoking an action whose flags you don't know. Use the command path without the binary name; " +
			"an empty command returns the root help listing all top-level commands. " +
			"For commands that take a JSON parameter (e.g. browser create --data) or when you need response fields and " +
			"error codes, use geelark_docs instead — it has much richer documentation.",
	}, s.handleHelp)
}

type helpArgs struct {
	Command string `json:"command" jsonschema:"command path without the binary name, e.g. \"phone automation add-task\"; empty for root help"`
}

var commandTokenRe = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9-]*$`)

func (s *Server) handleHelp(ctx context.Context, _ *mcp.CallToolRequest, args helpArgs) (*mcp.CallToolResult, any, error) {
	parts := strings.Fields(args.Command)
	for _, p := range parts {
		if !commandTokenRe.MatchString(p) {
			return errorResult("invalid command token: " + p), nil, nil
		}
	}
	oc := s.runHelp(ctx, parts)
	if !oc.OK() {
		return errorResult(oc.FailureText()), nil, nil
	}
	text := strings.TrimSpace(oc.Stdout)
	if text == "" {
		text = strings.TrimSpace(oc.Stderr)
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: text}},
	}, nil, nil
}
