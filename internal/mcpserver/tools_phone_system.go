package mcpserver

// geelark_phone_file — file & material library management.
// geelark_phone_system — device-level utilities: shell, ADB, webhook, OEM, analytics.
func phoneSystemTools() []*routedTool {
	file := []string{"phone", "file"}
	library := []string{"phone", "library"}
	auto := []string{"phone", "automation"}
	adb := []string{"phone", "adb"}
	shell := []string{"phone", "shell"}
	webhook := []string{"phone", "webhook"}
	oem := []string{"phone", "oem"}
	analytics := []string{"phone", "analytics"}

	return []*routedTool{
		{
			name: "geelark_phone_file",
			summary: "File and material management for cloud phones: upload files to phones or GeeLark temp storage, " +
				"query upload status, keybox upload, batch file upload, contact import, and the material Library " +
				"(search/create/delete materials and material tags).",
			actions: []action{
				act(file, "upload-to-phone", "Upload a file to a cloud phone"),
				act(file, "upload-temp", "Upload a local file to GeeLark temporary storage"),
				act(file, "upload-status", "Query the upload status of a file to the cloud phone"),
				act(file, "keybox-upload", "Upload a keybox file to a cloud phone"),
				act(file, "keybox-result", "Query the keybox upload task result"),
				act(auto, "file-upload", "Upload files to the cloud phone in batches (automation)"),
				act(auto, "import-contacts", "Batch import contacts to cloud phone (automation)"),
				act(library, "material-search", "Search materials in the Library"),
				act(library, "material-create", "Upload a file and create a material in the Library"),
				act(library, "material-delete", "Delete materials from the Library"),
				act(library, "tag-search", "Search material tags"),
				act(library, "tag-create", "Create a material tag"),
				act(library, "tag-delete", "Delete material tags"),
				act(library, "tag-set", "Set tags on materials"),
			},
		},
		{
			name: "geelark_phone_system",
			summary: "Cloud phone system utilities: execute shell commands, ADB connection management, webhook " +
				"callback URL, OEM/white-label customization, and analytics accounts & tags.",
			actions: []action{
				actAs(shell, "shell-exec", "exec", "Execute a shell command on running cloud phones (Android 10-16)"),
				actAs(adb, "adb-get-info", "get-info", "Get ADB connection information"),
				actAs(adb, "adb-set-status", "set-status", "Enable or disable ADB"),
				actAs(webhook, "webhook-get", "get", "Get webhook callback URL"),
				actAs(webhook, "webhook-set", "set", "Set webhook callback URL"),
				actAs(oem, "oem-customization", "customization", "Customize OEM/white label settings"),
				act(analytics, "accounts-list", "List analytics accounts"),
				act(analytics, "add-accounts", "Add analytics accounts (batch)"),
				act(analytics, "simple-add-account", "Quick add a single analytics account"),
				act(analytics, "update-account", "Update an analytics account"),
				act(analytics, "delete-account", "Delete an analytics account"),
				actAs(analytics, "analytics-data", "data", "Get analytics account data"),
				act(analytics, "tags-list", "List analytics tags"),
				act(analytics, "tags-create", "Create an analytics tag"),
				act(analytics, "tags-update", "Update an analytics tag"),
				act(analytics, "tags-delete", "Delete analytics tags"),
			},
		},
	}
}
