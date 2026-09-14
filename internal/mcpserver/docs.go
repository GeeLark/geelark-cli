package mcpserver

import (
	"context"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// registerDocs exposes the full reference documentation (flag tables, JSON
// parameter schemas, response fields, error codes, examples) for a command.
// Complements geelark_help: prefer geelark_docs for commands with JSON
// parameters (e.g. browser create --data).
func (s *Server) registerDocs(srv *mcp.Server) {
	mcp.AddTool(srv, &mcp.Tool{
		Name: "geelark_docs",
		Description: "Get the full reference documentation for a geelark-cli command: flag tables, JSON parameter " +
			"schemas (nested field tables), response fields, error codes, and examples. " +
			"Prefer this over geelark_help when the command takes a JSON parameter (e.g. browser create, proxy add) " +
			"or when you need response field/error code details. " +
			"Command path format: \"browser create\", \"phone automation tiktok-login\", \"proxy add\".",
	}, s.handleDocs)
}

type docsArgs struct {
	Command string `json:"command" jsonschema:"CLI command path, e.g. \"browser create\" or \"phone automation tiktok-login\" (leading geelark-cli is optional)"`
}

func (s *Server) handleDocs(_ context.Context, _ *mcp.CallToolRequest, args docsArgs) (*mcp.CallToolResult, any, error) {
	doc, err := s.docs.Get(args.Command)
	if err != nil {
		return errorResult(err.Error()), nil, nil
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: strings.TrimSpace(doc)}},
	}, nil, nil
}
