# phone automation multi-platform-video

Multichannel video distribution (TikTok / Instagram Reels / YouTube Shorts).

Use `--title` for a shared title across all platforms, or use per-platform titles (`--tiktok-title`, `--youtube-title`, `--instagram-title`) for individual titles. Use `--tiktok-recreate-link` / `--youtube-recreate-link` / `--instagram-recreate-link` to recreate the same style, along with `--same-style-voice` and `--original-voice`.

## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--video <csv>` | Comma-separated video URLs, max 10 (required) |
| `--title <text>` | Shared title for all platforms (max 100 chars) |
| `--tiktok-title <text>` | TikTok title (max 4000 chars) |
| `--youtube-title <text>` | YouTube title (max 100 chars) |
| `--instagram-title <text>` | Instagram title (max 2200 chars) |
| `--tiktok-recreate-link <text>` | TikTok style link |
| `--youtube-recreate-link <text>` | YouTube style link |
| `--instagram-recreate-link <text>` | Instagram style link |
| `--same-style-voice <n>` | Same style volume, 0-100 |
| `--original-voice <n>` | Original voice volume, 0-100 |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Shared title
geelark-cli phone automation multi-platform-video --id "557536075321468390" --schedule-at 1741846843 --title "title" --video "https://material.geelark.com/a.mp4"

# Per-platform titles with recreate link and voice volumes
geelark-cli phone automation multi-platform-video --id "557536075321468390" --schedule-at 1741846843 --tiktok-title "tt" --youtube-title "yt" --instagram-title "ig" --video "https://material.geelark.com/a.mp4" --tiktok-recreate-link "https://example.com" --same-style-voice 50 --original-voice 50
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
