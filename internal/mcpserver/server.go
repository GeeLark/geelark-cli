// Package mcpserver exposes geelark-cli over the Model Context Protocol.
//
// All GeeLark functionality is exposed through a small set of "routed" tools:
// each tool covers one module and dispatches on an `action` parameter that maps
// to a geelark-cli subcommand. Command flags are passed through the `params`
// object.
//
// Each tool call forks this same binary (os.Args[0]) so the child's stdout is
// isolated from the MCP transport. That keeps the existing CLI print sites
// unchanged and avoids version skew between MCP and CLI.
package mcpserver

import (
	"context"
	"io/fs"

	"github.com/geelark-tech/geelark-cli/internal/mcpdocs"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const instructions = "GeeLark cloud phone & browser management via geelark-cli. " +
	"Each tool dispatches on an 'action' parameter; pass command flags via the 'params' object " +
	"(keys are flag names without '--', values are string/number/boolean/array). " +
	"Use the geelark_help tool to get the full flag documentation for any action before calling it."

// Options configures a Server.
type Options struct {
	// SkillsFS holds the skills reference documentation tree.
	SkillsFS fs.FS
	// Version is reported to MCP clients during initialization.
	Version string
}

// Server serves the GeeLark MCP tools over stdio.
type Server struct {
	docs    *mcpdocs.Docs
	version string
}

// New creates a Server from opts.
func New(opts Options) *Server {
	return &Server{
		docs:    mcpdocs.New(opts.SkillsFS),
		version: opts.Version,
	}
}

// Run serves the MCP protocol over stdio until ctx is cancelled or the client
// disconnects. Nothing else may write to stdout for the lifetime of this call.
func (s *Server) Run(ctx context.Context) error {
	srv := mcp.NewServer(&mcp.Implementation{
		Name:    "geelark-mcp",
		Version: s.version,
	}, &mcp.ServerOptions{
		Instructions: instructions,
	})

	s.registerHelp(srv)
	s.registerDocs(srv)
	for _, t := range allTools() {
		s.registerRouted(srv, t)
	}

	return srv.Run(ctx, &mcp.StdioTransport{})
}
