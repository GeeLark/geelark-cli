package mcp

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"os"
	"os/signal"
	"syscall"

	"github.com/geelark-tech/geelark-cli/internal/mcpserver"
	"github.com/modelcontextprotocol/go-sdk/jsonrpc"
	"github.com/spf13/cobra"
)

// codeServerClosing is the JSON-RPC code the MCP SDK reports when the
// connection is torn down while the server is shutting down.
const codeServerClosing = -32004

const mcpLong = `Run geelark-cli as a Model Context Protocol (MCP) server over stdio.

AI agents (Cursor, Claude Code, VS Code, ...) talk to this server to manage
cloud phones, browsers, proxies, groups, tags and billing. Each tool call
forks this same binary, so the MCP tools always match the CLI surface and
the child's stdout never touches the MCP transport.

Authentication is shared with the CLI: run "geelark-cli config init" first,
or set credentials via environment variables (they override the config file):

    GEELARK_TOKEN            API token (with this set, no config file is needed)
    GEELARK_BASE_URL         Cloud Phone API base URL
    GEELARK_BROWSER_BASE_URL Browser API base URL

While the server is running stdout carries the MCP protocol and nothing else;
diagnostics go to stderr.

CLIENT CONFIGURATION:
    {
      "mcpServers": {
        "geelark": {
          "command": "npx",
          "args": ["-y", "geelark-cli", "mcp"],
          "env": {
            "GEELARK_TOKEN": "<your-token>",
            "GEELARK_BASE_URL": "https://openapi.geelark.com"
          }
        }
      }
    }`

// NewCmd creates the mcp command.
func NewCmd(skillsFS fs.FS, version string) *cobra.Command {
	return &cobra.Command{
		Use:           "mcp",
		Short:         "Run the MCP server over stdio",
		Long:          mcpLong,
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Signals are handled here rather than in the root command: the
			// server is the only command long-lived enough to need a graceful
			// shutdown, and intercepting signals globally would stop Ctrl+C
			// from killing the other commands, which never read this context.
			ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
			defer stop()

			srv := mcpserver.New(mcpserver.Options{
				SkillsFS: skillsFS,
				Version:  version,
			})
			if err := srv.Run(ctx); err != nil && !isShutdown(err) {
				return err
			}
			return nil
		},
	}
}

// isShutdown reports whether err is the client closing the pipe or a signal,
// both of which are normal ways for the server to stop.
//
// The SDK reports a closed connection as a jsonrpc.Error wrapping the
// underlying cause with %v, so the io.EOF check alone misses it.
func isShutdown(err error) bool {
	return errors.Is(err, io.EOF) ||
		errors.Is(err, context.Canceled) ||
		errors.Is(err, &jsonrpc.Error{Code: codeServerClosing})
}
