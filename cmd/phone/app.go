package phone

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newAppCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Application management",
		Long: `Manage applications — shop list, install/uninstall/start/stop apps on cloud phones,
upload apps, batch operations, and team app management.`,
	}

	cmd.AddCommand(newAppShopListCmd(newClient))
	cmd.AddCommand(newAppInstallCmd(newClient))
	cmd.AddCommand(newAppUninstallCmd(newClient))
	cmd.AddCommand(newAppStartCmd(newClient))
	cmd.AddCommand(newAppStopCmd(newClient))
	cmd.AddCommand(newAppListCmd(newClient))
	cmd.AddCommand(newAppInstallableListCmd(newClient))
	cmd.AddCommand(newAppUploadCmd(newClient))
	cmd.AddCommand(newAppUploadStatusCmd(newClient))
	cmd.AddCommand(newAppBatchCmd(newClient))
	cmd.AddCommand(newTeamAppCmd(newClient))

	return cmd
}

func newAppShopListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	var key string
	var getUploadApp bool

	cmd := &cobra.Command{
		Use:     "shop-list",
		Short:   "List applications from the app store",
		Long:    "Get the application list from the GeeLark app store.",
		Example: `  geelark-cli phone app shop-list --page 1 --page-size 10 --key "tiktok"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"page":     page,
				"pageSize": pageSize,
			}
			if key != "" {
				body["key"] = key
			}
			if getUploadApp {
				body["getUploadApp"] = true
			}

			result, err := c.PostAndPrint("/open/v1/app/shop/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of items per page (max 200)")
	cmd.Flags().StringVar(&key, "key", "", "Search keyword")
	cmd.Flags().BoolVar(&getUploadApp, "uploaded", false, "Get uploaded apps only")

	return cmd
}

func newAppInstallCmd(newClient clientFactory) *cobra.Command {
	var envID, appVersionID string

	cmd := &cobra.Command{
		Use:     "install",
		Short:   "Install an application on a cloud phone",
		Long:    "Install an app on a running cloud phone by environment ID and app version ID.",
		Example: `  geelark-cli phone app install --env-id "phone_id" --app-version-id "version_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"envId":        envID,
				"appVersionId": appVersionID,
			}

			result, err := c.PostAndPrint("/open/v1/app/install", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envID, "env-id", "", "Cloud phone environment ID (required)")
	cmd.Flags().StringVar(&appVersionID, "app-version-id", "", "App version ID (required)")
	_ = cmd.MarkFlagRequired("env-id")
	_ = cmd.MarkFlagRequired("app-version-id")

	return cmd
}

func newAppUninstallCmd(newClient clientFactory) *cobra.Command {
	var envID, packageName string

	cmd := &cobra.Command{
		Use:     "uninstall",
		Short:   "Uninstall an application from a cloud phone",
		Long:    "Uninstall an app from a running cloud phone by environment ID and package name.",
		Example: `  geelark-cli phone app uninstall --env-id "phone_id" --package-name "com.zhiliaoapp.musically"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"envId":       envID,
				"packageName": packageName,
			}

			result, err := c.PostAndPrint("/open/v1/app/uninstall", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envID, "env-id", "", "Cloud phone environment ID (required)")
	cmd.Flags().StringVar(&packageName, "package-name", "", "Application package name (required)")
	_ = cmd.MarkFlagRequired("env-id")
	_ = cmd.MarkFlagRequired("package-name")

	return cmd
}

func newAppStartCmd(newClient clientFactory) *cobra.Command {
	var envID, appVersionID, packageName string

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start an application on a cloud phone",
		Long: `Start an app on a running cloud phone.
Provide either --app-version-id or --package-name (package-name recommended).`,
		Example: `  geelark-cli phone app start --env-id "phone_id" --package-name "com.zhiliaoapp.musically"
  geelark-cli phone app start --env-id "phone_id" --app-version-id "version_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			if appVersionID == "" && packageName == "" {
				return fmt.Errorf("either --app-version-id or --package-name must be provided")
			}

			body := map[string]interface{}{
				"envId": envID,
			}
			if appVersionID != "" {
				body["appVersionId"] = appVersionID
			}
			if packageName != "" {
				body["packageName"] = packageName
			}

			result, err := c.PostAndPrint("/open/v1/app/start", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envID, "env-id", "", "Cloud phone environment ID (required)")
	cmd.Flags().StringVar(&appVersionID, "app-version-id", "", "App version ID")
	cmd.Flags().StringVar(&packageName, "package-name", "", "Application package name")
	_ = cmd.MarkFlagRequired("env-id")

	return cmd
}

func newAppStopCmd(newClient clientFactory) *cobra.Command {
	var envID, appVersionID, packageName string

	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop an application on a cloud phone",
		Long: `Stop (close) an app on a running cloud phone.
Provide either --app-version-id or --package-name (package-name recommended).`,
		Example: `  geelark-cli phone app stop --env-id "phone_id" --package-name "com.zhiliaoapp.musically"
  geelark-cli phone app stop --env-id "phone_id" --app-version-id "version_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			if appVersionID == "" && packageName == "" {
				return fmt.Errorf("either --app-version-id or --package-name must be provided")
			}

			body := map[string]interface{}{
				"envId": envID,
			}
			if appVersionID != "" {
				body["appVersionId"] = appVersionID
			}
			if packageName != "" {
				body["packageName"] = packageName
			}

			result, err := c.PostAndPrint("/open/v1/app/stop", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envID, "env-id", "", "Cloud phone environment ID (required)")
	cmd.Flags().StringVar(&appVersionID, "app-version-id", "", "App version ID")
	cmd.Flags().StringVar(&packageName, "package-name", "", "Application package name")
	_ = cmd.MarkFlagRequired("env-id")

	return cmd
}

func newAppListCmd(newClient clientFactory) *cobra.Command {
	var envID string
	var page, pageSize int

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List installed applications on a cloud phone",
		Long:    "Retrieve the list of applications installed on the specified cloud phone.",
		Example: `  geelark-cli phone app list --env-id "phone_id" --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"envId":    envID,
				"page":     page,
				"pageSize": pageSize,
			}

			result, err := c.PostAndPrint("/open/v1/app/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envID, "env-id", "", "Cloud phone environment ID (required)")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of items per page (max 100)")
	_ = cmd.MarkFlagRequired("env-id")

	return cmd
}

func newAppInstallableListCmd(newClient clientFactory) *cobra.Command {
	var envID string
	var page, pageSize int
	var name string
	var getUploadApp bool

	cmd := &cobra.Command{
		Use:   "installable-list",
		Short: "List applications available for installation on a cloud phone",
		Long:  "Get the list of apps available for installation on the specified cloud phone.",
		Example: `  geelark-cli phone app installable-list --env-id "phone_id" --page 1 --page-size 10
  geelark-cli phone app installable-list --env-id "phone_id" --name "tiktok"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"envId":    envID,
				"page":     page,
				"pageSize": pageSize,
			}
			if name != "" {
				body["name"] = name
			}
			if getUploadApp {
				body["getUploadApp"] = true
			}

			result, err := c.PostAndPrint("/open/v1/app/installable/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envID, "env-id", "", "Cloud phone environment ID (required)")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of items per page (max 100)")
	cmd.Flags().StringVar(&name, "name", "", "Search keyword")
	cmd.Flags().BoolVar(&getUploadApp, "uploaded", false, "Get uploaded apps only")
	_ = cmd.MarkFlagRequired("env-id")

	return cmd
}

func newAppUploadCmd(newClient clientFactory) *cobra.Command {
	var fileURL, desc string

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload an application (APK/XAPK)",
		Long: `Upload an application to the team. Supports APK and XAPK files only.
Provide a file URL (e.g. from temp upload or library).`,
		Example: `  geelark-cli phone app upload --url "https://material.geelark.cn/xxx.apk"
  geelark-cli phone app upload --url "https://material.geelark.cn/xxx.apk" --desc "my app"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"fileUrl": fileURL,
			}
			if desc != "" {
				body["desc"] = desc
			}

			result, err := c.PostAndPrint("/open/v1/app/upload", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&fileURL, "url", "", "Application file URL (APK/XAPK) (required)")
	cmd.Flags().StringVar(&desc, "desc", "", "Description/remark")
	_ = cmd.MarkFlagRequired("url")

	return cmd
}

func newAppUploadStatusCmd(newClient clientFactory) *cobra.Command {
	var taskID string

	cmd := &cobra.Command{
		Use:     "upload-status",
		Short:   "Query application upload status",
		Long:    "Query the upload status of an application upload task.",
		Example: `  geelark-cli phone app upload-status --task-id "1830906144634757120"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"taskId": taskID,
			}

			result, err := c.PostAndPrint("/open/v1/app/upload/status", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&taskID, "task-id", "", "Upload task ID (required)")
	_ = cmd.MarkFlagRequired("task-id")

	return cmd
}

func newAppBatchCmd(newClient clientFactory) *cobra.Command {
	var action int
	var groupIds, packageName, versionID string

	cmd := &cobra.Command{
		Use:   "batch",
		Short: "Batch operate apps on cloud phones",
		Long: `Batch operation of applications on opened cloud phones.
Actions: 1=Start, 2=Stop, 3=Restart, 4=Install, 5=Uninstall.
For start/stop/restart/uninstall, provide --package-name.
For install, provide --version-id.
Optionally filter by --group-ids.`,
		Example: `  geelark-cli phone app batch --action 1 --package-name "com.zhiliaoapp.musically"
  geelark-cli phone app batch --action 4 --version-id "1793552962140770305" --group-ids "group1,group2"
  geelark-cli phone app batch --action 5 --package-name "com.zhiliaoapp.musically"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			if packageName == "" && versionID == "" {
				return fmt.Errorf("either --package-name or --version-id must be provided")
			}

			body := map[string]interface{}{
				"action": action,
			}
			if groupIds != "" {
				body["groupIds"] = strings.Split(groupIds, ",")
			}
			if packageName != "" {
				body["packageName"] = packageName
			}
			if versionID != "" {
				body["versionId"] = versionID
			}

			result, err := c.PostAndPrint("/open/v1/app/operation/batch", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&action, "action", 0, "Operation: 1=Start, 2=Stop, 3=Restart, 4=Install, 5=Uninstall (required)")
	cmd.Flags().StringVar(&groupIds, "group-ids", "", "Comma-separated group IDs (default: all groups)")
	cmd.Flags().StringVar(&packageName, "package-name", "", "Application package name (for start/stop/restart/uninstall)")
	cmd.Flags().StringVar(&versionID, "version-id", "", "Application version ID (for install)")
	_ = cmd.MarkFlagRequired("action")

	return cmd
}

func newTeamAppCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "team-app",
		Short: "Team application management",
		Long:  "Manage team applications — list, add, remove, set authorization/auto-start/keep-alive/ROOT/auto-install.",
	}

	cmd.AddCommand(newTeamAppListCmd(newClient))
	cmd.AddCommand(newTeamAppAddCmd(newClient))
	cmd.AddCommand(newTeamAppRemoveCmd(newClient))
	cmd.AddCommand(newTeamAppSetAuthCmd(newClient))
	cmd.AddCommand(newTeamAppSetAutoStartCmd(newClient))
	cmd.AddCommand(newTeamAppSetKeepAliveCmd(newClient))
	cmd.AddCommand(newTeamAppSetRootCmd(newClient))
	cmd.AddCommand(newTeamAppSetAutoInstallCmd(newClient))

	return cmd
}

func newTeamAppListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List team applications",
		Long:    "Get the list of team applications.",
		Example: `  geelark-cli phone app team-app list --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"page":     page,
				"pageSize": pageSize,
			}

			result, err := c.PostAndPrint("/open/v1/app/teamApp/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of items per page (max 200)")

	return cmd
}

func newTeamAppAddCmd(newClient clientFactory) *cobra.Command {
	var id, versionID string
	var installGroupIds string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add an app to team applications",
		Long: `Add the app to the team applications. It will be automatically installed after the cloud phone starts.
Use --install-group-ids to restrict which environment groups can install the app.`,
		Example: `  geelark-cli phone app team-app add --id "1793552962123993090" --version-id "1793552962140770305"
  geelark-cli phone app team-app add --id "1793552962123993090" --version-id "1793552962140770305" --install-group-ids "group1,group2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id":        id,
				"versionId": versionID,
			}
			if installGroupIds != "" {
				body["installGroupIds"] = strings.Split(installGroupIds, ",")
			}

			result, err := c.PostAndPrint("/open/v1/app/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Application ID (required)")
	cmd.Flags().StringVar(&versionID, "version-id", "", "Version ID (required)")
	cmd.Flags().StringVar(&installGroupIds, "install-group-ids", "", "Comma-separated environment group IDs (default: all groups; \"0\" for ungrouped)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("version-id")

	return cmd
}

func newTeamAppRemoveCmd(newClient clientFactory) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:     "remove",
		Short:   "Remove an app from team applications",
		Long:    "Remove the application from the team applications.",
		Example: `  geelark-cli phone app team-app remove --id "497652752864775437"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id": id,
			}

			result, err := c.PostAndPrint("/open/v1/app/remove", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Team application ID (required)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newTeamAppSetAuthCmd(newClient clientFactory) *cobra.Command {
	var id string
	var appAuth int

	cmd := &cobra.Command{
		Use:   "set-auth",
		Short: "Set team app authorization",
		Long: `Grant or revoke team application permissions (including location permissions, etc.).
Only applies to newly installed applications.`,
		Example: `  geelark-cli phone app team-app set-auth --id "497652752864775437" --auth 1
  geelark-cli phone app team-app set-auth --id "497652752864775437" --auth 0`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id":      id,
				"appAuth": appAuth,
			}

			result, err := c.PostAndPrint("/open/v1/app/auth/status", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Team application ID (required)")
	cmd.Flags().IntVar(&appAuth, "auth", 0, "Authorization: 0=disable, 1=enable (required)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newTeamAppSetAutoStartCmd(newClient clientFactory) *cobra.Command {
	var id string
	var opera int

	cmd := &cobra.Command{
		Use:     "set-auto-start",
		Short:   "Enable or disable team app auto-start",
		Long:    "Enable or disable auto-start for a team application.",
		Example: `  geelark-cli phone app team-app set-auto-start --id "497652752864775437" --opera 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id":    id,
				"opera": opera,
			}

			result, err := c.PostAndPrint("/open/v1/app/setAutoStart", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Team application ID (required)")
	cmd.Flags().IntVar(&opera, "opera", 0, "Operation: 0=disable, 1=enable (required)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newTeamAppSetKeepAliveCmd(newClient clientFactory) *cobra.Command {
	var id string
	var opera int

	cmd := &cobra.Command{
		Use:   "set-keep-alive",
		Short: "Enable or disable team app keep-alive",
		Long: `Enable or disable keep-alive for a team application. Requires Pro plan.
Only Android 12/13/14/15 supported. One app can be kept alive at max.
On Android 12/13/15, restart the app after enabling. On Android 14, it takes effect immediately.`,
		Example: `  geelark-cli phone app team-app set-keep-alive --id "497652752864775437" --opera 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id":    id,
				"opera": opera,
			}

			result, err := c.PostAndPrint("/open/v1/app/setKeepAlive", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Team application ID (required)")
	cmd.Flags().IntVar(&opera, "opera", 0, "Operation: 0=disable, 1=enable (required)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newTeamAppSetRootCmd(newClient clientFactory) *cobra.Command {
	var id string
	var opera int

	cmd := &cobra.Command{
		Use:     "set-root",
		Short:   "Enable or disable team app ROOT access",
		Long:    "Enable or disable ROOT access for a team application.",
		Example: `  geelark-cli phone app team-app set-root --id "497652752864775437" --opera 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id":    id,
				"opera": opera,
			}

			result, err := c.PostAndPrint("/open/v1/app/root", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Team application ID (required)")
	cmd.Flags().IntVar(&opera, "opera", 0, "Operation: 0=disable, 1=enable (required)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newTeamAppSetAutoInstallCmd(newClient clientFactory) *cobra.Command {
	var id string
	var status int
	var installGroupIds string

	cmd := &cobra.Command{
		Use:   "set-auto-install",
		Short: "Set team app auto-install",
		Long: `Enable or disable automatic installation for a team application.
When enabled, the app will be installed after the cloud phone starts up.
Use --install-group-ids to restrict which environment groups can install the app.`,
		Example: `  geelark-cli phone app team-app set-auto-install --id "497652752864775437" --status 1
  geelark-cli phone app team-app set-auto-install --id "497652752864775437" --status 1 --install-group-ids "group1,group2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id":     id,
				"status": status,
			}
			if installGroupIds != "" {
				body["installGroupIds"] = strings.Split(installGroupIds, ",")
			}

			result, err := c.PostAndPrint("/open/v1/app/setStatus", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Team application ID (required)")
	cmd.Flags().IntVar(&status, "status", 0, "Auto-install: 0=disable, 1=enable (required)")
	cmd.Flags().StringVar(&installGroupIds, "install-group-ids", "", "Comma-separated environment group IDs (default: all groups; \"0\" for ungrouped)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}
