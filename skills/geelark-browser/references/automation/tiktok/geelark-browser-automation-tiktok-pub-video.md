# browser automation tiktok-pub-video

Create a TikTok automation task to publish a video with optional caption, cover, location and AI label.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--video-path <text>` | Local video file path on your computer (max 999 chars, required) |
| `--caption <text>` | Video caption (max 4000 chars) |
| `--cover-path <text>` | Local cover image file path on your computer (max 999 chars) |
| `--location <text>` | Location (max 999 chars) |
| `--ai-label` | Enable AI label |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Publish a TikTok video
geelark-cli browser automation tiktok-pub-video --eid "557536075321468390" --schedule-at 1741846843 --video-path "C:\video\demo.mp4"

# Publish with caption and AI label
geelark-cli browser automation tiktok-pub-video --eid "557536075321468390" --schedule-at 1741846843 --video-path "C:\video\demo.mp4" --caption "hello tiktok #funny" --ai-label
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
