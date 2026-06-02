# geelark-browser automation

Browser automation task management. Create, query, cancel, and retry RPA tasks across platforms (TikTok, YouTube, X/Twitter, Instagram, Facebook, Reddit) and custom task flows.

## Command Overview

| Command | Description |
|---------|-------------|
| `task-search` | Query browser task list |
| `task-detail` | Query browser task details with logs |
| `task-cancel` | Cancel waiting/in-progress tasks |
| `task-restart` | Retry failed/cancelled tasks |
| `add-custom-task` | Create custom automation task |
| `task-flow` | Query browser custom task flows |
| `cookie-bot` | Create Cookie Bot task |
| `tiktok-search` | TikTok search videos, likes and comments |
| `tiktok-comment` | TikTok like and comment on videos |
| `tiktok-like` | TikTok like specified videos |
| `youtube-watch` | YouTube watch videos |
| `x-post` | X (Twitter) post a tweet |
| `x-like` | X (Twitter) like and retweet tweets |
| `instagram-search` | Instagram search hashtags and browse posts |
| `instagram-like` | Browse and like Instagram feed |
| `facebook-post` | Facebook post a status |
| `facebook-homepage` | Facebook account creates a homepage |
| `facebook-friends` | Facebook add recommended friends |
| `facebook-like` | Facebook like all on the first screen |
| `reddit-like` | Browse and like Reddit posts searched by keywords |

## Common Flags

All automation task creation commands share these common optional flags:

| Flag | Description |
|------|-------------|
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

All automation task creation commands require:

| Flag | Description |
|------|-------------|
| `--eid <text>` | Browser environment ID |
| `--schedule-at <n>` | Schedule time, second-level timestamp |

## Task Status

| Status | Value | Description |
|--------|-------|-------------|
| Waiting | 1 | Task is queued |
| In progress | 2 | Task is running |
| Completed | 3 | Task finished successfully |
| Failed | 4 | Task failed |
| Cancelled | 7 | Task was cancelled |

## Task Failure Codes

| Code | Reason |
|------|--------|
| 20001 | Task start time is earlier than current time |
| 20002 | Browser failed to start |
| 29999 | Task was interrupted |

## References

### Task Management
- [task-search](task/geelark-browser-automation-task-search.md)
- [task-detail](task/geelark-browser-automation-task-detail.md)
- [task-cancel](task/geelark-browser-automation-task-cancel.md)
- [task-restart](task/geelark-browser-automation-task-restart.md)

### Custom Task
- [add-custom-task](custom-task/geelark-browser-automation-add-custom-task.md)
- [task-flow](custom-task/geelark-browser-automation-task-flow.md)

### Other
- [cookie-bot](other/geelark-browser-automation-cookie-bot.md)

### TikTok
- [tiktok-search](tiktok/geelark-browser-automation-tiktok-search.md)
- [tiktok-comment](tiktok/geelark-browser-automation-tiktok-comment.md)
- [tiktok-like](tiktok/geelark-browser-automation-tiktok-like.md)

### YouTube
- [youtube-watch](youtube/geelark-browser-automation-youtube-watch.md)

### X (Twitter)
- [x-post](x/geelark-browser-automation-x-post.md)
- [x-like](x/geelark-browser-automation-x-like.md)

### Instagram
- [instagram-search](instagram/geelark-browser-automation-instagram-search.md)
- [instagram-like](instagram/geelark-browser-automation-instagram-like.md)

### Facebook
- [facebook-post](facebook/geelark-browser-automation-facebook-post.md)
- [facebook-homepage](facebook/geelark-browser-automation-facebook-homepage.md)
- [facebook-friends](facebook/geelark-browser-automation-facebook-friends.md)
- [facebook-like](facebook/geelark-browser-automation-facebook-like.md)

### Reddit
- [reddit-like](reddit/geelark-browser-automation-reddit-like.md)
