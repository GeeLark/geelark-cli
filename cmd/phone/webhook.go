package phone

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newWebhookCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Webhook management",
		Long:  "Manage webhook callback URLs — get and set the callback URL.",
	}

	cmd.AddCommand(newWebhookGetCmd(newClient))
	cmd.AddCommand(newWebhookSetCmd(newClient))

	return cmd
}

func newWebhookGetCmd(newClient clientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "get",
		Short:   "Get webhook callback URL",
		Long:    "Retrieve the currently set webhook callback URL.",
		Example: `  geelark-cli phone webhook get`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			result, err := c.PostAndPrint("/open/v1/callback/get", nil)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func newWebhookSetCmd(newClient clientFactory) *cobra.Command {
	var url string

	cmd := &cobra.Command{
		Use:     "set",
		Short:   "Set webhook callback URL",
		Long:    "Set the webhook callback URL for receiving notifications.",
		Example: `  geelark-cli phone webhook set --url "https://example.com/callback"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"url": url,
			}

			result, err := c.PostAndPrint("/open/v1/callback/set", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&url, "url", "", "Callback URL (required)")
	_ = cmd.MarkFlagRequired("url")

	return cmd
}
