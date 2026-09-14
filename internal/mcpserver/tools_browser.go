package mcpserver

// geelark_browser — browser lifecycle management.
// geelark_browser_automation — browser automation tasks & task management.
func browserTools() []*routedTool {
	browser := []string{"browser"}
	auto := []string{"browser", "automation"}

	return []*routedTool{
		{
			name: "geelark_browser",
			summary: "Manage GeeLark browsers via the local API (requires the GeeLark client running and logged in): " +
				"list, create, edit, start, stop, delete, clone, clear cache, bookmarks, cookies, kernels, transfer.",
			actions: []action{
				act(browser, "list", "List browsers"),
				act(browser, "create", "Create a new browser"),
				act(browser, "simple-create", "Quick create a single browser"),
				act(browser, "edit", "Edit a browser"),
				act(browser, "start", "Launch a browser"),
				act(browser, "stop", "Close a browser"),
				act(browser, "delete", "Delete browsers"),
				act(browser, "clone", "Clone a browser"),
				act(browser, "clear-cache", "Clear browser cache"),
				act(browser, "check-status", "Check browser startup status"),
				act(browser, "api-status", "Check API interface availability"),
				act(browser, "get-bookmark", "Get browser bookmarks"),
				act(browser, "set-bookmark", "Set browser bookmarks"),
				act(browser, "get-cookie", "Get browser cookies"),
				act(browser, "get-kernels", "List available browser kernels"),
				act(browser, "update-kernels", "Download and update a browser kernel"),
				act(browser, "move-group", "Move browsers to a group"),
				act(browser, "transfer", "Transfer browsers to another account"),
				act(browser, "ext-group-list", "List browser extension categories"),
			},
		},
		{
			name: "geelark_browser_automation",
			summary: "Browser automation tasks (geelark-cli browser automation ...): social platform tasks for " +
				"Facebook, Instagram, TikTok, X (Twitter), Reddit, YouTube, plus Cookie Bot, custom tasks, and task management.",
			actions: []action{
				act(auto, "task-search", "Query browser task list"),
				act(auto, "task-detail", "Query browser task details"),
				act(auto, "task-cancel", "Cancel a browser task"),
				act(auto, "task-restart", "Retry a browser task"),
				act(auto, "task-flow", "Query browser custom task flows"),
				act(auto, "add-custom-task", "Create a custom automation task"),
				act(auto, "cookie-bot", "Create a Cookie Bot task"),
				// Facebook
				act(auto, "facebook-post", "Facebook post a status"),
				act(auto, "facebook-friends", "Facebook add recommended friends"),
				act(auto, "facebook-homepage", "Facebook account creates a homepage"),
				act(auto, "facebook-like", "Facebook like all on the first screen"),
				// Instagram
				act(auto, "instagram-like", "Browse and like Instagram feed"),
				act(auto, "instagram-search", "Instagram search hashtags and browse posts"),
				// TikTok
				act(auto, "tiktok-like", "TikTok like specified videos"),
				act(auto, "tiktok-comment", "TikTok like and comment on videos"),
				act(auto, "tiktok-search", "TikTok search videos, likes and comments"),
				// X (Twitter)
				act(auto, "x-like", "X (Twitter) like and retweet tweets"),
				act(auto, "x-post", "X (Twitter) retweet and post a tweet"),
				// Reddit
				act(auto, "reddit-like", "Browse and like Reddit posts searched by keywords"),
				// YouTube
				act(auto, "youtube-watch", "YouTube watch videos"),
			},
		},
	}
}
