# phone automation facebook-pub-reels

Facebook publish Reels video.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--description <text>` | Caption (max 500 chars, required) |
| `--video <text>` | Video URL (required) |
| `--page <text>` | Page |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation facebook-pub-reels --id "557536075321468390" --schedule-at 1741846843 --description "My reel" --video "https://material.geelark.com/a.mp4"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
