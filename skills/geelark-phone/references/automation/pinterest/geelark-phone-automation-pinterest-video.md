# phone automation pinterest-video

Publish video on Pinterest.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--title <text>` | Title (max 100 chars, required) |
| `--description <text>` | Description (max 800 chars, required) |
| `--video <csv>` | Comma-separated video URLs (required) |
| `--link <text>` | Link |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation pinterest-video --id "557536075321468390" --schedule-at 1741846843 --title "title" --description "desc" --video "https://material.geelark.com/a.mp4"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
