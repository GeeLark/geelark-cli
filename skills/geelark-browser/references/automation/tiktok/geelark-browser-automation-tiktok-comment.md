# browser automation tiktok-comment

Create a TikTok automation task to like and comment on videos with specified comments.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--comments <csv>` | Comma-separated comments (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Create TikTok comment task
geelark-cli browser automation tiktok-comment --eid "557536075321468390" --schedule-at 1741846843 --comments "hello,great"

# With name
geelark-cli browser automation tiktok-comment --eid "557536075321468390" --schedule-at 1741846843 --comments "hello" --name "myTask"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
