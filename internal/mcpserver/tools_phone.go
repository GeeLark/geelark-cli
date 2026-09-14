package mcpserver

// geelark_phone — cloud phone lifecycle management (geelark-cli phone ...).
// geelark_phone_app — application management (geelark-cli phone app ...).
func phoneTools() []*routedTool {
	phone := []string{"phone"}
	app := []string{"phone", "app"}
	teamApp := []string{"phone", "app", "team-app"}

	return []*routedTool{
		{
			name: "geelark_phone",
			summary: "Manage GeeLark cloud phones: list, create, start, stop, restart, delete, clone, update, " +
				"screenshot, GPS, SMS, network config, transfer, live streaming push, and more.",
			actions: []action{
				act(phone, "list", "List all cloud phones"),
				act(phone, "status", "Query cloud phone status"),
				act(phone, "create", "Create new cloud phones"),
				act(phone, "simple-create", "Quick create a single cloud phone"),
				act(phone, "start", "Start cloud phones"),
				act(phone, "stop", "Stop cloud phones"),
				act(phone, "restart", "Restart a cloud phone"),
				act(phone, "delete", "Delete cloud phones"),
				act(phone, "clone", "Clone a cloud phone"),
				act(phone, "update", "Update cloud phone information (name, remark, group)"),
				act(phone, "move-group", "Move cloud phones to a group"),
				act(phone, "transfer", "Transfer cloud phones to another account"),
				act(phone, "new-one", "One-click new machine (reset cloud phone identity)"),
				act(phone, "screenshot", "Take a screenshot of a cloud phone (async, then query screenshot-result)"),
				act(phone, "screenshot-result", "Get screenshot task result"),
				act(phone, "get-device-id", "Get cloud phone device ID"),
				act(phone, "get-gps", "Get GPS information of cloud phones"),
				act(phone, "set-gps", "Set GPS for cloud phones"),
				act(phone, "send-sms", "Send SMS to a cloud phone"),
				act(phone, "set-root", "Set root status on cloud phones"),
				act(phone, "hide-accessibility", "Hide accessibility in apps"),
				act(phone, "net-config-get", "Get cloud phone network config"),
				act(phone, "net-config-set", "Set cloud phone network config"),
				act(phone, "set-net-type", "Set cloud phone network type"),
				act(phone, "brand-list", "List cloud phone brands and models"),
				act(phone, "brand-team-list", "List team-uploaded cloud phone brands and models"),
				act(phone, "import-contacts", "Import contacts to a cloud phone"),
				act(phone, "import-contacts-result", "Get import contacts task result"),
				act(phone, "video-push", "Upload a video and start live streaming push"),
				act(phone, "video-push-result", "Get live streaming push task result"),
				act(phone, "video-push-stop", "Stop live streaming push"),
			},
		},
		{
			name: "geelark_phone_app",
			summary: "Manage applications on cloud phones (geelark-cli phone app ...): install/uninstall/start/stop apps, " +
				"app store listings, APK upload, batch operations, and team app settings.",
			actions: []action{
				act(app, "list", "List installed applications on a cloud phone"),
				act(app, "installable-list", "List applications available for installation on a cloud phone"),
				act(app, "shop-list", "List applications from the app store"),
				act(app, "install", "Install an application on a cloud phone"),
				act(app, "uninstall", "Uninstall an application from a cloud phone"),
				act(app, "start", "Start an application on a cloud phone"),
				act(app, "stop", "Stop an application on a cloud phone"),
				act(app, "batch", "Batch operate apps on cloud phones"),
				act(app, "upload", "Upload an application (APK/XAPK)"),
				act(app, "upload-status", "Query application upload status"),
				actAs(teamApp, "team-app-add", "add", "Add an app to team applications"),
				actAs(teamApp, "team-app-list", "list", "List team applications"),
				actAs(teamApp, "team-app-remove", "remove", "Remove an app from team applications"),
				actAs(teamApp, "team-app-set-auth", "set-auth", "Set team app authorization"),
				actAs(teamApp, "team-app-set-auto-install", "set-auto-install", "Set team app auto-install"),
				actAs(teamApp, "team-app-set-auto-start", "set-auto-start", "Enable or disable team app auto-start"),
				actAs(teamApp, "team-app-set-hide", "set-hide", "Set team app hidden status"),
				actAs(teamApp, "team-app-set-keep-alive", "set-keep-alive", "Enable or disable team app keep-alive"),
				actAs(teamApp, "team-app-set-root", "set-root", "Enable or disable team app ROOT access"),
			},
		},
	}
}
