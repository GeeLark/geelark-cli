package browser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newAutomationCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "automation",
		Short: "Browser automation tasks",
		Long: `Manage browser automation tasks via the local API.
Includes task management, custom tasks, and platform-specific automation
(TikTok, YouTube, X/Twitter, Instagram, Facebook, Reddit, Cookie Bot).`,
	}

	cmd.AddCommand(newTaskSearchCmd(newClient))
	cmd.AddCommand(newTaskDetailCmd(newClient))
	cmd.AddCommand(newTaskCancelCmd(newClient))
	cmd.AddCommand(newTaskRestartCmd(newClient))
	cmd.AddCommand(newCustomTaskCmd(newClient))
	cmd.AddCommand(newTaskFlowCmd(newClient))
	cmd.AddCommand(newCookieBotCmd(newClient))
	cmd.AddCommand(newTikTokSearchCmd(newClient))
	cmd.AddCommand(newTikTokCommentCmd(newClient))
	cmd.AddCommand(newTikTokLikeCmd(newClient))
	cmd.AddCommand(newYouTubeWatchCmd(newClient))
	cmd.AddCommand(newXPostCmd(newClient))
	cmd.AddCommand(newXLikeCmd(newClient))
	cmd.AddCommand(newInstagramSearchCmd(newClient))
	cmd.AddCommand(newInstagramLikeCmd(newClient))
	cmd.AddCommand(newFacebookPostCmd(newClient))
	cmd.AddCommand(newFacebookHomepageCmd(newClient))
	cmd.AddCommand(newFacebookFriendsCmd(newClient))
	cmd.AddCommand(newFacebookLikeCmd(newClient))
	cmd.AddCommand(newRedditLikeCmd(newClient))

	return cmd
}

func newTaskSearchCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	var taskIds string

	cmd := &cobra.Command{
		Use:     "task-search",
		Short:   "Query browser task list",
		Long:    "Get the list of browser automation tasks with optional filters.",
		Example: `  geelark-cli browser automation task-search --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{}
			if page > 0 {
				body["page"] = page
			}
			if pageSize > 0 {
				body["pageSize"] = pageSize
			}
			if taskIds != "" {
				body["taskIds"] = strings.Split(taskIds, ",")
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/search", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Page size (max 100)")
	cmd.Flags().StringVar(&taskIds, "task-ids", "", "Comma-separated task IDs")

	return cmd
}

func newTaskDetailCmd(newClient clientFactory) *cobra.Command {
	var taskId string

	cmd := &cobra.Command{
		Use:     "task-detail",
		Short:   "Query browser task details",
		Long:    "Get detailed information of a browser automation task by task ID.",
		Example: `  geelark-cli browser automation task-detail --task-id "497652752864775437"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"taskId": taskId}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/detail", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&taskId, "task-id", "", "Task ID (required)")
	_ = cmd.MarkFlagRequired("task-id")
	return cmd
}

func newTaskCancelCmd(newClient clientFactory) *cobra.Command {
	var taskId string

	cmd := &cobra.Command{
		Use:     "task-cancel",
		Short:   "Cancel a browser task",
		Long:    "Cancel a browser automation task. Can cancel while running or waiting to be executed.",
		Example: `  geelark-cli browser automation task-cancel --task-id "497652752864775437"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"taskId": taskId}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/cancel", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&taskId, "task-id", "", "Task ID (required)")
	_ = cmd.MarkFlagRequired("task-id")
	return cmd
}

func newTaskRestartCmd(newClient clientFactory) *cobra.Command {
	var taskId string

	cmd := &cobra.Command{
		Use:     "task-restart",
		Short:   "Retry a browser task",
		Long:    "Retry a browser automation task. Can retry if the task failed or was cancelled.",
		Example: `  geelark-cli browser automation task-restart --task-id "497652752864775437"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"taskId": taskId}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/restart", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&taskId, "task-id", "", "Task ID (required)")
	_ = cmd.MarkFlagRequired("task-id")
	return cmd
}

func newCustomTaskCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var flowID, paramMapJSON string

	cmd := &cobra.Command{
		Use:   "add-custom-task",
		Short: "Create a custom automation task",
		Long: `Create a custom browser automation task using a task flow ID.
First call 'task-flow' to get available task flows, then create a task with the flow ID.`,
		Example: `  geelark-cli browser automation add-custom-task --eid "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885" --param-map "{\"Title\":\"video\",\"Desc\":\"this is video\"}"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
				"flowId":     flowID,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if paramMapJSON != "" {
				var paramMap interface{}
				if err := json.Unmarshal([]byte(paramMapJSON), &paramMap); err != nil {
					return fmt.Errorf("invalid --param-map JSON: %w", err)
				}
				body["paramMap"] = paramMap
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&flowID, "flow-id", "", "Task flow ID (required)")
	cmd.Flags().StringVar(&paramMapJSON, "param-map", "", "Task flow parameters as JSON string")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("flow-id")
	return cmd
}

func newTaskFlowCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int

	cmd := &cobra.Command{
		Use:     "task-flow",
		Short:   "Query browser custom task flows",
		Long:    "Get the list of available custom task flows for browser automation.",
		Example: `  geelark-cli browser automation task-flow --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{}
			if page > 0 {
				body["page"] = page
			}
			if pageSize > 0 {
				body["pageSize"] = pageSize
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/flow", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Page size (max 100)")
	return cmd
}

func newCookieBotCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var pages string

	cmd := &cobra.Command{
		Use:   "cookie-bot",
		Short: "Create a Cookie Bot task",
		Long:  "Create a Cookie Bot automation task that visits specified webpages to collect cookies.",
		Example: `  geelark-cli browser automation cookie-bot --eid "557536075321468390" --schedule-at 1741846843 --pages "https://a.com,https://b.com"
  geelark-cli browser automation cookie-bot --eid "557536075321468390" --schedule-at 1741846843 --pages "https://a.com" --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
				"pages":      strings.Split(pages, ","),
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/cookieBot", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&pages, "pages", "", "Comma-separated webpage URLs to visit (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("pages")
	return cmd
}

func newTikTokSearchCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var searchKeyword string

	cmd := &cobra.Command{
		Use:   "tiktok-search",
		Short: "TikTok search videos, likes and comments",
		Long:  "Create a TikTok automation task to search videos by keyword, like and comment on them.",
		Example: `  geelark-cli browser automation tiktok-search --eid "557536075321468390" --schedule-at 1741846843 --search-keyword "hello"
  geelark-cli browser automation tiktok-search --eid "557536075321468390" --schedule-at 1741846843 --search-keyword "hello" --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":           eid,
				"scheduleAt":    scheduleAt,
				"searchKeyword": searchKeyword,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/tiktokSearch", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&searchKeyword, "search-keyword", "", "Search keyword (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("search-keyword")
	return cmd
}

func newTikTokCommentCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var comments string

	cmd := &cobra.Command{
		Use:   "tiktok-comment",
		Short: "TikTok like and comment on videos",
		Long:  "Create a TikTok automation task to like and comment on videos with specified comments.",
		Example: `  geelark-cli browser automation tiktok-comment --eid "557536075321468390" --schedule-at 1741846843 --comments "hello,great"
  geelark-cli browser automation tiktok-comment --eid "557536075321468390" --schedule-at 1741846843 --comments "hello" --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
				"comments":   strings.Split(comments, ","),
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/tiktokComment", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&comments, "comments", "", "Comma-separated comments (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("comments")
	return cmd
}

func newTikTokLikeCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var videoLink, comment string

	cmd := &cobra.Command{
		Use:   "tiktok-like",
		Short: "TikTok like specified videos",
		Long:  "Create a TikTok automation task to like a specified video, optionally with a comment.",
		Example: `  geelark-cli browser automation tiktok-like --eid "557536075321468390" --schedule-at 1741846843 --video-link "https://www.tiktok.com/video/38210380122"
  geelark-cli browser automation tiktok-like --eid "557536075321468390" --schedule-at 1741846843 --video-link "https://www.tiktok.com/video/38210380122" --comment "nice"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
				"videoLink":  videoLink,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if comment != "" {
				body["comment"] = comment
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/tiktokLike", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&videoLink, "video-link", "", "TikTok video link (required)")
	cmd.Flags().StringVar(&comment, "comment", "", "Comment on the video")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("video-link")
	return cmd
}

func newYouTubeWatchCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var searchKeyword, title, comment string

	cmd := &cobra.Command{
		Use:     "youtube-watch",
		Short:   "YouTube watch videos",
		Long:    "Create a YouTube automation task to search and watch videos, with optional title and comment.",
		Example: `  geelark-cli browser automation youtube-watch --eid "557536075321468390" --schedule-at 1741846843 --search-keyword "hello" --title "myTitle" --comment "myComment"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":           eid,
				"scheduleAt":    scheduleAt,
				"searchKeyword": searchKeyword,
				"title":         title,
				"comment":       comment,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/youtubeWatch", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&searchKeyword, "search-keyword", "", "Search keyword (required)")
	cmd.Flags().StringVar(&title, "title", "", "Video title (required)")
	cmd.Flags().StringVar(&comment, "comment", "", "Comment (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("search-keyword")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("comment")
	return cmd
}

func newXPostCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var content string

	cmd := &cobra.Command{
		Use:   "x-post",
		Short: "X (Twitter) retweet and post a tweet",
		Long:  "Create an X (Twitter) automation task to post a tweet with specified content.",
		Example: `  geelark-cli browser automation x-post --eid "557536075321468390" --schedule-at 1741846843 --content "hello"
  geelark-cli browser automation x-post --eid "557536075321468390" --schedule-at 1741846843 --content "hello" --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
				"content":    content,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/xPost", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&content, "content", "", "Tweet content (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("content")
	return cmd
}

func newXLikeCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64

	cmd := &cobra.Command{
		Use:   "x-like",
		Short: "X (Twitter) like and retweet tweets",
		Long:  "Create an X (Twitter) automation task to like and retweet tweets in the feed.",
		Example: `  geelark-cli browser automation x-like --eid "557536075321468390" --schedule-at 1741846843
  geelark-cli browser automation x-like --eid "557536075321468390" --schedule-at 1741846843 --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/xLike", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newInstagramSearchCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var searchKeywords string

	cmd := &cobra.Command{
		Use:   "instagram-search",
		Short: "Instagram search hashtags and browse posts",
		Long:  "Create an Instagram automation task to search hashtags and browse posts.",
		Example: `  geelark-cli browser automation instagram-search --eid "557536075321468390" --schedule-at 1741846843 --search-keywords "hello,world"
  geelark-cli browser automation instagram-search --eid "557536075321468390" --schedule-at 1741846843 --search-keywords "hello" --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":            eid,
				"scheduleAt":     scheduleAt,
				"searchKeywords": strings.Split(searchKeywords, ","),
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/instagramSearch", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&searchKeywords, "search-keywords", "", "Comma-separated search keywords (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("search-keywords")
	return cmd
}

func newInstagramLikeCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64

	cmd := &cobra.Command{
		Use:   "instagram-like",
		Short: "Browse and like Instagram feed",
		Long:  "Create an Instagram automation task to browse and like posts in the feed.",
		Example: `  geelark-cli browser automation instagram-like --eid "557536075321468390" --schedule-at 1741846843
  geelark-cli browser automation instagram-like --eid "557536075321468390" --schedule-at 1741846843 --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/instagramLike", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newFacebookPostCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var content string

	cmd := &cobra.Command{
		Use:   "facebook-post",
		Short: "Facebook post a status",
		Long:  "Create a Facebook automation task to post a status with specified content.",
		Example: `  geelark-cli browser automation facebook-post --eid "557536075321468390" --schedule-at 1741846843 --content "hello"
  geelark-cli browser automation facebook-post --eid "557536075321468390" --schedule-at 1741846843 --content "hello" --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
				"content":    content,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/facebookPost", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&content, "content", "", "Status content (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("content")
	return cmd
}

func newFacebookHomepageCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var mainPageName string
	var category string

	cmd := &cobra.Command{
		Use:   "facebook-homepage",
		Short: "Facebook account creates a homepage",
		Long:  "Create a Facebook automation task to create a homepage with specified name and category.",
		Example: `  geelark-cli browser automation facebook-homepage --eid "557536075321468390" --schedule-at 1741846843 --main-page-name "myPage" --category "dev"
  geelark-cli browser automation facebook-homepage --eid "557536075321468390" --schedule-at 1741846843 --main-page-name "myPage" --category "dev,business"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":          eid,
				"scheduleAt":   scheduleAt,
				"mainPageName": mainPageName,
				"category":     strings.Split(category, ","),
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/facebookHomepage", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&mainPageName, "main-page-name", "", "Homepage name (required)")
	cmd.Flags().StringVar(&category, "category", "", "Comma-separated homepage categories (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("main-page-name")
	_ = cmd.MarkFlagRequired("category")
	return cmd
}

func newFacebookFriendsCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64

	cmd := &cobra.Command{
		Use:   "facebook-friends",
		Short: "Facebook add recommended friends",
		Long:  "Create a Facebook automation task to add recommended friends.",
		Example: `  geelark-cli browser automation facebook-friends --eid "557536075321468390" --schedule-at 1741846843
  geelark-cli browser automation facebook-friends --eid "557536075321468390" --schedule-at 1741846843 --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/facebookFriends", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newFacebookLikeCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64

	cmd := &cobra.Command{
		Use:   "facebook-like",
		Short: "Facebook like all on the first screen",
		Long:  "Create a Facebook automation task to like all posts on the first screen.",
		Example: `  geelark-cli browser automation facebook-like --eid "557536075321468390" --schedule-at 1741846843
  geelark-cli browser automation facebook-like --eid "557536075321468390" --schedule-at 1741846843 --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":        eid,
				"scheduleAt": scheduleAt,
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/facebookLike", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newRedditLikeCmd(newClient clientFactory) *cobra.Command {
	var eid, name, remark string
	var scheduleAt int64
	var searchKeywords string

	cmd := &cobra.Command{
		Use:   "reddit-like",
		Short: "Browse and like Reddit posts searched by keywords",
		Long:  "Create a Reddit automation task to browse and like posts found by keyword search.",
		Example: `  geelark-cli browser automation reddit-like --eid "557536075321468390" --schedule-at 1741846843 --search-keywords "hello,world"
  geelark-cli browser automation reddit-like --eid "557536075321468390" --schedule-at 1741846843 --search-keywords "hello" --name "myTask"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"eid":            eid,
				"scheduleAt":     scheduleAt,
				"searchKeywords": strings.Split(searchKeywords, ","),
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/task/redditLike", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&eid, "eid", "", "Environment ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&searchKeywords, "search-keywords", "", "Comma-separated search keywords (required)")
	_ = cmd.MarkFlagRequired("eid")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("search-keywords")
	return cmd
}
