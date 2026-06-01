package phone

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newADBCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "adb",
		Short: "ADB management",
		Long:  "Manage ADB connections — get connection info and enable/disable ADB on cloud phones.",
	}

	cmd.AddCommand(newADBGetInfoCmd(newClient))
	cmd.AddCommand(newADBSetStatusCmd(newClient))

	return cmd
}

func newADBGetInfoCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "get-info",
		Short:   "Get ADB connection information",
		Long:    "Retrieve ADB connection info (IP, port, password) for cloud phones.",
		Example: `  geelark-cli phone adb get-info --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}

			result, err := c.PostAndPrint("/open/v1/adb/getData", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs (required)")
	_ = cmd.MarkFlagRequired("ids")

	return cmd
}

func newADBSetStatusCmd(newClient clientFactory) *cobra.Command {
	var ids string
	var open bool

	cmd := &cobra.Command{
		Use:   "set-status",
		Short: "Enable or disable ADB",
		Long: `Enable or disable ADB on cloud phones.
Supports Android 9/11/12/13/14/15/16. Cloud phone must be started first.
After enabling, wait ~3 seconds before retrieving connection info.`,
		Example: `  geelark-cli phone adb set-status --ids "id1,id2" --open
  geelark-cli phone adb set-status --ids "id1" --open=false`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids":  strings.Split(ids, ","),
				"open": open,
			}

			result, err := c.PostAndPrint("/open/v1/adb/setStatus", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs (required)")
	cmd.Flags().BoolVar(&open, "open", true, "Enable (true) or disable (false) ADB")
	_ = cmd.MarkFlagRequired("ids")

	return cmd
}
