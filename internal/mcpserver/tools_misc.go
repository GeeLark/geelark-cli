package mcpserver

// geelark_proxy / geelark_group / geelark_tag / geelark_billing / geelark_account.
func miscTools() []*routedTool {
	proxy := []string{"proxy"}
	group := []string{"group"}
	tag := []string{"tag"}
	billing := []string{"billing"}
	auth := []string{"auth"}
	config := []string{"config"}

	return []*routedTool{
		{
			name:    "geelark_proxy",
			summary: "Proxy management: add/update/delete/list proxies, proxy detection, and proxy group management.",
			actions: []action{
				act(proxy, "list", "List all proxies"),
				act(proxy, "add", "Add proxies (batch)"),
				act(proxy, "simple-add", "Quick add a single proxy"),
				act(proxy, "update", "Update proxies (batch)"),
				act(proxy, "simple-update", "Quick update a single proxy"),
				act(proxy, "delete", "Delete proxies"),
				act(proxy, "check", "Check/detect a proxy"),
				act(proxy, "group-list", "List proxy groups"),
				act(proxy, "group-add", "Create a proxy group"),
				act(proxy, "group-update", "Update a proxy group name"),
				act(proxy, "group-delete", "Delete proxy groups"),
			},
		},
		{
			name:    "geelark_group",
			summary: "Cloud phone & browser group management: create, update, delete, list groups.",
			actions: []action{
				act(group, "list", "List groups"),
				act(group, "create", "Create groups"),
				act(group, "simple-create", "Quick create a single group"),
				act(group, "update", "Update groups"),
				act(group, "simple-update", "Quick update a single group"),
				act(group, "delete", "Delete groups"),
			},
		},
		{
			name:    "geelark_tag",
			summary: "Cloud phone & browser tag management: create, update, delete, list tags.",
			actions: []action{
				act(tag, "list", "List tags"),
				act(tag, "create", "Create tags"),
				act(tag, "simple-create", "Quick create a single tag"),
				act(tag, "update", "Update tags"),
				act(tag, "simple-update", "Quick update a single tag"),
				act(tag, "delete", "Delete tags"),
			},
		},
		{
			name:    "geelark_billing",
			summary: "Billing and subscription management: balance, transactions, plan info/list/renew/upgrade, time add-ons.",
			actions: []action{
				act(billing, "balance", "Query account balance"),
				act(billing, "transaction-detail", "Query billing transaction details"),
				act(billing, "plan-info", "Get current subscription plan info"),
				act(billing, "plan-list", "Get all available plans"),
				act(billing, "plan-renew", "Renew subscription plan"),
				act(billing, "plan-upgrade", "Upgrade subscription plan"),
				act(billing, "buy-time-addon", "Buy time add-on minutes"),
			},
		},
		{
			name: "geelark_account",
			summary: "CLI account & configuration: initialize credentials (token + API base URLs), check " +
				"authentication status, show current configuration.",
			actions: []action{
				actAs(config, "config-init", "init", "Initialize CLI credentials: pass --token (required), "+
					"--base-url (Cloud Phone API, default https://openapi.geelark.com), "+
					"--browser-base-url (default http://localhost:40185)"),
				actAs(auth, "auth-status", "status", "Check authentication status"),
				actAs(config, "config-show", "show", "Show current CLI configuration"),
			},
		},
	}
}
