package mcp

import (
	"context"
	"errors"
	"io"
	"io/fs"

	"github.com/geelark-tech/geelark-cli/internal/mcpserver"
	"github.com/spf13/cobra"
)

const mcpLong = `Run geelark-cli as a Model Context Protocol (MCP) server over stdio.

AI agents (Cursor, Claude Code, VS Code, ...) talk to this server to manage
cloud phones, browsers, proxies, groups, tags and billing. Each tool call
forks this same binary, so the MCP tools always match the CLI surface and
the child's stdout never touches the MCP transport.

Authentication is shared with the CLI: run "geelark-cli config init" first.

While the server is running stdout carries the MCP protocol and nothing else;
diagnostics go to stderr.

CLIENT CONFIGURATION:
    {
      "mcpServers": {
        "geelark": {
          "command": "npx",
          "args": ["-y", "geelark-cli", "mcp"]
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
			srv := mcpserver.New(mcpserver.Options{
				SkillsFS: skillsFS,
				Version:  version,
			})
			err := srv.Run(cmd.Context())
			// A client closing the pipe, or a signal, is a normal shutdown.
			if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
				return nil
			}
			return err
		},
	}
}
