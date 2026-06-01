package auth

import (
	"encoding/json"
	"fmt"

	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/geelark-tech/geelark-cli/internal/output"
	"github.com/spf13/cobra"
)

type clientFactory func() (*client.Client, error)

func NewCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication management",
		Long:  "Check authentication status and verify API token validity.",
	}

	cmd.AddCommand(newStatusCmd(newClient))

	return cmd
}

func newStatusCmd(newClient clientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "status",
		Short:   "Check authentication status",
		Long:    "Verify the current API token. Returns standard JSON envelope.",
		Example: `  geelark-cli auth status`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			resp, err := c.Post("/open/v1/auth/status", map[string]interface{}{})
			if err != nil {
				return fmt.Errorf("authentication check failed: %w", err)
			}

			result, err := json.MarshalIndent(resp, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(output.FormatResponse(string(result)))
			return nil
		},
	}
}
