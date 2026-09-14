package mcpserver

// Cloud phone automation tools (geelark-cli phone automation ...),
// split into TikTok, other social platforms, and task management.
func phoneAutomationTools() []*routedTool {
	auto := []string{"phone", "automation"}

	return []*routedTool{
		{
			name: "geelark_phone_automation_tiktok",
			summary: "TikTok automation tasks on cloud phones (geelark-cli phone automation ...): login, AI comments, " +
				"likes (star), follows, private messages, video/image publishing, hiding and deleting videos and comments. " +
				"'-asia' variants use the Asia-region channel.",
			actions: []action{
				act(auto, "add-task", "Add a TikTok video/image/warmup task (publish video, image set, or warm up account)"),
				act(auto, "tiktok-login", "TikTok account login"),
				act(auto, "tiktok-edit-profile", "TikTok profile edit"),
				act(auto, "tiktok-comment", "TikTok AI comment"),
				act(auto, "tiktok-comment-asia", "TikTok AI comment (Asia)"),
				act(auto, "tiktok-star", "TikTok random star (like)"),
				act(auto, "tiktok-star-asia", "TikTok random star (like) (Asia)"),
				act(auto, "tiktok-follow", "TikTok random follow"),
				act(auto, "tiktok-follow-asia", "TikTok random follow (Asia)"),
				act(auto, "tiktok-message", "Send private message on TikTok"),
				act(auto, "tiktok-message-asia", "Send private message on TikTok (Asia)"),
				act(auto, "tiktok-hide", "Hide TikTok videos"),
				act(auto, "tiktok-hide-asia", "Hide TikTok videos (Asia)"),
				act(auto, "tiktok-delete", "Delete all TikTok videos"),
				act(auto, "tiktok-delete-asia", "Delete all TikTok videos (Asia)"),
				act(auto, "tiktok-delete-comment", "Delete TikTok comments"),
				act(auto, "tiktok-delete-comment-asia", "Delete TikTok comments (Asia)"),
			},
		},
		{
			name: "geelark_phone_automation_social",
			summary: "Social platform automation tasks on cloud phones (geelark-cli phone automation ...): Facebook, " +
				"Instagram, YouTube, X (Twitter), Reddit, Threads, Pinterest, Google, SHEIN — publish content, " +
				"auto login, warm up accounts, send private messages, edit profiles.",
			actions: []action{
				// Facebook
				act(auto, "facebook-publish", "Facebook post content"),
				act(auto, "facebook-pub-reels", "Facebook publish Reels video"),
				act(auto, "facebook-auto-comment", "Facebook auto comment"),
				act(auto, "facebook-login", "Facebook auto login"),
				act(auto, "facebook-maintenance", "Facebook account maintenance"),
				act(auto, "facebook-message", "Send private message on Facebook"),
				act(auto, "facebook-reels-active", "Facebook Reels maintenance"),
				// Instagram
				act(auto, "instagram-pub-reels", "Instagram publish Reels video"),
				act(auto, "instagram-pub-reels-images", "Instagram publish Reels image"),
				act(auto, "instagram-ai-comment", "Instagram AI random comment"),
				act(auto, "instagram-login", "Instagram auto login"),
				act(auto, "instagram-warmup", "Instagram AI account warmup"),
				act(auto, "instagram-message", "Send private message on Instagram"),
				act(auto, "instagram-follow-account", "Follow Instagram accounts"),
				act(auto, "instagram-edit-profile", "Edit Instagram profile"),
				// YouTube
				act(auto, "youtube-pub-video", "YouTube publish video"),
				act(auto, "youtube-pub-short", "YouTube publish Short"),
				act(auto, "youtube-maintenance", "YouTube account maintenance"),
				act(auto, "youtube-edit-profile", "Edit YouTube profile"),
				// X (Twitter)
				act(auto, "x-publish", "Publish content on X (Twitter)"),
				// Reddit
				act(auto, "reddit-video", "Publish video on Reddit"),
				act(auto, "reddit-image", "Publish pictures and texts on Reddit"),
				act(auto, "reddit-warmup", "Reddit AI account warmup"),
				// Threads
				act(auto, "threads-video", "Publish video on Threads"),
				act(auto, "threads-image", "Publish pictures and texts on Threads"),
				// Pinterest
				act(auto, "pinterest-video", "Publish video on Pinterest"),
				act(auto, "pinterest-image", "Publish pictures and texts on Pinterest"),
				// Google
				act(auto, "google-login", "Google auto login"),
				act(auto, "google-app-download", "Download apps on Google Play"),
				act(auto, "google-app-browser", "Open the app on Google for browsing"),
				// SHEIN
				act(auto, "shein-login", "SHEIN auto login"),
				// Multi-platform
				act(auto, "multi-platform-video", "Multichannel video distribution (TikTok/Instagram Reels/YouTube Shorts)"),
			},
		},
		{
			name: "geelark_phone_task",
			summary: "Cloud phone automation task management (geelark-cli phone automation ...): query, cancel, retry " +
				"tasks, and create/import/export custom task flows.",
			actions: []action{
				act(auto, "task-query", "Query cloud phone tasks by IDs"),
				act(auto, "task-history", "Batch query task history"),
				act(auto, "task-detail", "Query cloud phone task detail"),
				act(auto, "task-cancel", "Cancel cloud phone tasks"),
				act(auto, "task-restart", "Retry cloud phone tasks"),
				act(auto, "add-custom-task", "Create a custom automation task"),
				act(auto, "task-flow-list", "Query custom task flows"),
				act(auto, "task-flow-import", "Import or update a custom task flow"),
				act(auto, "task-flow-export", "Export a custom task flow"),
			},
		},
	}
}
