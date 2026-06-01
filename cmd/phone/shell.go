package phone

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newShellCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shell",
		Short: "Execute shell commands on cloud phones",
		Long:  "Execute shell commands on running cloud phones. Supports Android 10/12/13/14/15/16.",
	}

	cmd.AddCommand(newShellExecCmd(newClient))

	return cmd
}

func newShellExecCmd(newClient clientFactory) *cobra.Command {
	var id, cmdStr string

	cmd := &cobra.Command{
		Use:   "exec",
		Short: "Execute a shell command",
		Long: `Execute a shell command on a running cloud phone.
Only supports Android 10/12/13/14/15/16 models.`,
		Example: `  geelark-cli phone shell exec --id "phone_id" --cmd "pm list packages"
  geelark-cli phone shell exec --id "phone_id" --cmd "ls /sdcard/Download"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id":  id,
				"cmd": cmdStr,
			}

			result, err := c.PostAndPrint("/open/v1/shell/execute", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&cmdStr, "cmd", "", "Shell command to execute (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("cmd")

	return cmd
}
