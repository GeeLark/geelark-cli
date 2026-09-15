package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

Settings are merged into the existing configuration: flags you omit keep their
current value, so "config init --base-url ..." changes only that endpoint.

Two API endpoints are configurable:
- Cloud Phone API (default: https://openapi.geelark.com) — used by phone, proxy, group, tag, billing, app, shell, adb, webhook, analytics, oem
- Browser API (default: http://localhost:40185, local) — used by browser commands. Token is shared.`,
		Example: `  geelark-cli config init --token "your_api_token"
  geelark-cli config init --base-url "https://openapi.geelark.com"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Start from the stored settings so that omitting a flag keeps its
			// current value instead of resetting it to the default.
			cfg := loadOrEmpty()

			// Asking questions is the zero-flag setup flow. Once any setting is
			// passed the run is scripted, so the other values are kept as they
			// are rather than blocking on stdin.
			interactive := !cmd.Flags().Changed("token") &&
				!cmd.Flags().Changed("base-url") &&
				!cmd.Flags().Changed("browser-base-url")

			// A single reader for every question: bufio reads ahead, so a second
			// reader on os.Stdin would drop the lines the first one buffered.
			in := bufio.NewReader(os.Stdin)

			if cmd.Flags().Changed("token") {
				cfg.Token = token
			} else if interactive {
				t, err := promptToken(in, cfg.Token != "")
				if err != nil {
					return err
				}
				if t != "" {
					cfg.Token = t
				}
			}

			if cmd.Flags().Changed("base-url") {
				cfg.BaseURL = baseURL
			} else if interactive {
				v, err := promptEndpoint(in, "Cloud Phone API base URL", cfg.BaseURL, config.DefaultBaseURL)
				if err != nil {
					return err
				}
				if v != "" {
					cfg.BaseURL = v
				}
			}

			if cmd.Flags().Changed("browser-base-url") {
				cfg.BrowserBaseURL = browserBaseURL
			} else if interactive {
				v, err := promptEndpoint(in, "Browser API base URL", cfg.BrowserBaseURL, config.DefaultBrowserBaseURL)
				if err != nil {
					return err
				}
				if v != "" {
					cfg.BrowserBaseURL = v
				}
			}

			if cfg.Token == "" {
				return fmt.Errorf("token cannot be empty")
			}
			clearDefaultEndpoints(cfg)

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

// loadOrEmpty returns the settings stored on disk, or an empty set when there
// is nothing usable to merge into. A file that cannot be read or parsed is
// reported and then ignored rather than being an error, because init is also
// how a broken config gets repaired.
//
// The file is read directly instead of through config.Load so that unset
// endpoints stay empty: Load substitutes the defaults, which would then be
// written back and pin the user to today's values.
func loadOrEmpty() *config.Config {
	path := config.ConfigFile()

	b, err := os.ReadFile(path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return &config.Config{}
	case err != nil:
		fmt.Fprintf(os.Stderr, "Warning: ignoring unreadable config %s: %v\n", path, err)
		return &config.Config{}
	}

	var cfg config.Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: ignoring malformed config %s: %v\n", path, err)
		return &config.Config{}
	}
	return &cfg
}

// clearDefaultEndpoints drops endpoints that match the built-in defaults, which
// config.Load also substitutes when they are unset. Keeping them out of the
// file means a user who never pinned an endpoint follows the default if it
// ever changes.
func clearDefaultEndpoints(cfg *config.Config) {
	if cfg.BaseURL == config.DefaultBaseURL {
		cfg.BaseURL = ""
	}
	if cfg.BrowserBaseURL == config.DefaultBrowserBaseURL {
		cfg.BrowserBaseURL = ""
	}
}

// promptToken reads a token from in, returning "" to keep the current one.
// Prompts go to stderr and are only printed on a terminal, so stdout stays
// clean for piped and non-interactive callers — the MCP server runs the CLI
// without stdin — and the captured output is not polluted.
func promptToken(in *bufio.Reader, hasCurrent bool) (string, error) {
	if stdinIsTerminal() {
		if hasCurrent {
			fmt.Fprint(os.Stderr, "Enter your GeeLark API token (leave empty to keep the current one): ")
		} else {
			fmt.Fprint(os.Stderr, "Enter your GeeLark API token: ")
		}
	}

	t, err := in.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("failed to read token: %w", err)
	}
	return strings.TrimSpace(t), nil
}

// promptEndpoint reads an API endpoint from in, returning "" to keep the
// current value. Mirrors promptToken: the prompt is printed to stderr only on
// a terminal, but in is always read, so piped multi-line input works while
// the MCP server (no stdin) just sees EOF and keeps the current value.
// Note stdinIsTerminal also reports true for /dev/null (a character device),
// which the stderr routing keeps harmless.
func promptEndpoint(in *bufio.Reader, label, current, def string) (string, error) {
	if stdinIsTerminal() {
		display := current
		if display == "" {
			display = def
		}
		fmt.Fprintf(os.Stderr, "%s (enter to keep %s): ", label, display)
	}

	v, err := in.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("failed to read %s: %w", label, err)
	}
	return strings.TrimSpace(v), nil
}

func stdinIsTerminal() bool {
	fi, err := os.Stdin.Stat()
	return err == nil && fi.Mode()&os.ModeCharDevice != 0
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
