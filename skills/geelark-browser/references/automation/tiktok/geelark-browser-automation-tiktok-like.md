# browser automation tiktok-like

Create a TikTok automation task to like a specified video, optionally with a comment.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--video-link <text>` | TikTok video link (required) |
| `--comment <text>` | Comment on the video |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Like a TikTok video
geelark-cli browser automation tiktok-like --eid "557536075321468390" --schedule-at 1741846843 --video-link "https://www.tiktok.com/video/38210380122"

# Like with comment
geelark-cli browser automation tiktok-like --eid "557536075321468390" --schedule-at 1741846843 --video-link "https://www.tiktok.com/video/38210380122" --comment "nice"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
