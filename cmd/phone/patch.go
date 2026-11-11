package phone

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newPatchCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "patch",
		Short: "Xposed patch management",
		Long: `Manage Xposed patches on cloud phones. Supports Android 12/13/14/15/16.
Android 12/13/15 use the original install mode; Android 14/16 use the apmt install mode and require extra patchName/targetApp parameters on install.`,
	}

	cmd.AddCommand(newPatchInstallCmd(newClient))
	cmd.AddCommand(newPatchUninstallCmd(newClient))
	cmd.AddCommand(newPatchStatusCmd(newClient))

	return cmd
}

func newPatchInstallCmd(newClient clientFactory) *cobra.Command {
	var id, apkURL, patchName, targetApp string

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install an Xposed patch",
		Long: `Install an Xposed patch on a cloud phone.
Android 12/13/15: apkURL supports APK and XAPK; patchName/targetApp are not needed.
Android 14/16: apkURL supports APK only; patchName and targetApp are required.`,
		Example: `  geelark-cli phone patch install --id "493188646250219877" --apk-url "https://material.geelark.com/xxx.apk"
  geelark-cli phone patch install --id "493188646250219877" --apk-url "https://material.geelark.com/xxx.apk" --patch-name "geelarkpatch" --target-app "com.abc"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"id":     id,
				"apkURL": apkURL,
			}
			if patchName != "" {
				body["patchName"] = patchName
			}
			if targetApp != "" {
				body["targetApp"] = targetApp
			}
			result, err := c.PostAndPrint("/open/v1/phone/patch/install", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&apkURL, "apk-url", "", "Patch install file URL (required)")
	cmd.Flags().StringVar(&patchName, "patch-name", "", "Patch name, 1-64 chars, letters/digits/underscore/hyphen only, no leading hyphen (required for Android 14/16)")
	cmd.Flags().StringVar(&targetApp, "target-app", "", "Hook target app package name, max 128 chars; 'system' means system hook (required for Android 14/16)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("apk-url")

	return cmd
}

func newPatchUninstallCmd(newClient clientFactory) *cobra.Command {
	var id, patchId, patchName string

	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall an Xposed patch",
		Long: `Uninstall an Xposed patch from a cloud phone.
Android 12/13/15: pass --patch-id.
Android 14/16: pass --patch-name (must match the one used at install time).`,
		Example: `  geelark-cli phone patch uninstall --id "532930887177274368" --patch-id "sdktest"
  geelark-cli phone patch uninstall --id "643654639442503208" --patch-name "hookplugin-debug"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"id": id,
			}
			if patchId != "" {
				body["patchId"] = patchId
			}
			if patchName != "" {
				body["patchName"] = patchName
			}
			result, err := c.PostAndPrint("/open/v1/phone/patch/uninstall", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&patchId, "patch-id", "", "Patch ID, for Android 12/13/15")
	cmd.Flags().StringVar(&patchName, "patch-name", "", "Patch name used at install time, for Android 14/16")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newPatchStatusCmd(newClient clientFactory) *cobra.Command {
	var patchId, id, taskId string
	var queryType int

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Query Xposed patch install status",
		Long: `Query the install status of an Xposed patch.
type 1: query by patchId (Android 12/13/15 only, requires --patch-id and --id).
type 2: query by taskId (all supported Android versions, requires --task-id).`,
		Example: `  geelark-cli phone patch status --type 2 --task-id "1830906144634757120"
  geelark-cli phone patch status --type 1 --patch-id "sdktest" --id "532930887177274368"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"type": queryType,
			}
			if patchId != "" {
				body["patchId"] = patchId
			}
			if id != "" {
				body["id"] = id
			}
			if taskId != "" {
				body["taskId"] = taskId
			}
			result, err := c.PostAndPrint("/open/v1/phone/patch/install/status", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&queryType, "type", 0, "Query type: 1=by patchId (Android 12/13/15), 2=by taskId (all) (required)")
	cmd.Flags().StringVar(&patchId, "patch-id", "", "Patch ID, required when type=1")
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID, required when type=1")
	cmd.Flags().StringVar(&taskId, "task-id", "", "Task ID returned by patch install, required when type=2")
	_ = cmd.MarkFlagRequired("type")

	return cmd
}
