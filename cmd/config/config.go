package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/geelark-tech/geelark-cli/internal/config"
	"github.com/geelark-tech/geelark-cli/internal/output"
	"github.com/spf13/cobra"
)

// NewCmd creates the config command group.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage GeeLark CLI configuration",
		Long:  "Configure and view GeeLark CLI credentials and settings.",
	}

	cmd.AddCommand(newInitCmd())
	cmd.AddCommand(newShowCmd())

	return cmd
}

func newInitCmd() *cobra.Command {
	var token string
	var baseURL string
	var browserBaseURL string

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize GeeLark CLI configuration",
		Long: `Set up GeeLark CLI credentials interactively or via flags.

Get your token from the GeeLark client settings page.

Two API endpoints are configurable:
- Cloud Phone API (default: https://openapi.geelark.com) — used by phone, proxy, group, tag, billing, app, shell, adb, webhook, analytics, oem
- Browser API (default: http://localhost:40185, local) — used by browser commands. Token is shared.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if token == "" {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter your GeeLark API token: ")
				t, err := reader.ReadString('\n')
				if err != nil {
					return fmt.Errorf("failed to read token: %w", err)
				}
				token = strings.TrimSpace(t)
			}

			if token == "" {
				return fmt.Errorf("token cannot be empty")
			}

			cfg := &config.Config{
				Token:          token,
				BaseURL:        baseURL,
				BrowserBaseURL: browserBaseURL,
			}

			if err := config.Save(cfg); err != nil {
				return err
			}

			output.PrintSuccess(fmt.Sprintf("Configuration saved to %s", config.ConfigFile()))
			return nil
		},
	}

	cmd.Flags().StringVar(&token, "token", "", "GeeLark API token (shared by Cloud Phone and Browser APIs)")
	cmd.Flags().StringVar(&baseURL, "base-url", "", "Cloud Phone API base URL (default: https://openapi.geelark.com)")
	cmd.Flags().StringVar(&browserBaseURL, "browser-base-url", "", "Browser API base URL (default: http://localhost:40185)")

	return cmd
}

func newShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}

			fmt.Printf("Config file:        %s\n", config.ConfigFile())
			fmt.Printf("Cloud Phone API:    %s\n", cfg.BaseURL)
			fmt.Printf("Browser API:        %s\n", cfg.BrowserBaseURL)
			// Mask token for security
			masked := cfg.Token
			if len(masked) > 8 {
				masked = masked[:4] + "****" + masked[len(masked)-4:]
			}
			fmt.Printf("Token:              %s\n", masked)
			return nil
		},
	}
}
