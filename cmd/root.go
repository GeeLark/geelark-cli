package cmd

import (
	"fmt"
	"os"

	"github.com/geelark-tech/geelark-cli/cmd/auth"
	"github.com/geelark-tech/geelark-cli/cmd/billing"
	"github.com/geelark-tech/geelark-cli/cmd/browser"
	cmdconfig "github.com/geelark-tech/geelark-cli/cmd/config"
	"github.com/geelark-tech/geelark-cli/cmd/group"
	"github.com/geelark-tech/geelark-cli/cmd/phone"
	"github.com/geelark-tech/geelark-cli/cmd/proxy"
	"github.com/geelark-tech/geelark-cli/cmd/tag"
	"github.com/geelark-tech/geelark-cli/internal/build"
	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/geelark-tech/geelark-cli/internal/config"
	"github.com/geelark-tech/geelark-cli/internal/output"
	"github.com/spf13/cobra"
)

const rootLong = `geelark-cli — GeeLark Cloud Phone & Browser CLI tool.

USAGE:
    geelark-cli <command> [subcommand] [flags]

EXAMPLES:
    # List all cloud phones
    geelark-cli phone list --page 1 --page-size 10

    # List browsers (local API)
    geelark-cli browser list --page 1 --page-size 10

    # Start cloud phones
    geelark-cli phone start --ids "id1,id2,id3"

    # Check balance
    geelark-cli billing balance

GLOBAL FLAGS:
    --format <fmt>    Output format: json (default) | pretty | table

COMMANDS:
    config      Configure credentials (init, show)
    auth        Authentication management (status)
    phone       Cloud phone management (includes adb, shell, file, webhook, oem, analytics, app)
    browser     Browser management (local API)
    proxy       Proxy management
    group       Group management
    tag         Tag management
    billing     Billing & subscription management

DOCS:
    https://open.geelark.com

More help: geelark-cli <command> --help`

// Execute runs the root command and returns the process exit code.
func Execute() int {
	var formatFlag string

	rootCmd := &cobra.Command{
		Use:           "geelark-cli",
		Short:         "GeeLark Cloud Phone & Browser CLI tool",
		Long:          rootLong,
		Version:       build.Version,
		SilenceUsage:  false,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if formatFlag != "" {
				output.Format = formatFlag
			}
		},
	}

	rootCmd.PersistentFlags().StringVar(&formatFlag, "format", "json", "Output format: json | pretty | table")

	newClient := func() (*client.Client, error) {
		cfg, err := config.Load()
		if err != nil {
			return nil, err
		}
		return client.New(cfg), nil
	}

	// Register sub-commands
	rootCmd.AddCommand(cmdconfig.NewCmd())
	rootCmd.AddCommand(auth.NewCmd(newClient))
	rootCmd.AddCommand(phone.NewCmd(newClient))
	rootCmd.AddCommand(browser.NewCmd(newClient))
	rootCmd.AddCommand(proxy.NewCmd(newClient))
	rootCmd.AddCommand(group.NewCmd(newClient))
	rootCmd.AddCommand(tag.NewCmd(newClient))
	rootCmd.AddCommand(billing.NewCmd(newClient))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return 1
	}
	return 0
}
