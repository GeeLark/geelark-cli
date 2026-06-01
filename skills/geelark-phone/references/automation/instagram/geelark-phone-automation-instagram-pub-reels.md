# phone automation instagram-pub-reels

Instagram publish Reels video.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--description <text>` | Caption (max 2200 chars, required) |
| `--video <csv>` | Comma-separated video URLs, max 10 (required) |
| `--same-style-url <text>` | Same style URL |
| `--same-style-voice <n>` | Same style volume, 0-100 |
| `--original-voice <n>` | Original volume, 0-100 |
| `--ai-tag` | AI tag (default false) |
| `--need-share-link` | Whether to retrieve sharing link (default false) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation instagram-pub-reels --id "557536075321468390" --schedule-at 1741846843 --description "My reel" --video "https://material.geelark.com/a.mp4"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
