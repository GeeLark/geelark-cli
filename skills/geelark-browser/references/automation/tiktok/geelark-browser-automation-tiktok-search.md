# browser automation tiktok-search

Create a TikTok automation task to search videos by keyword, like and comment on them.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--search-keyword <text>` | Search keyword (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Create TikTok search task
geelark-cli browser automation tiktok-search --eid "557536075321468390" --schedule-at 1741846843 --search-keyword "hello"

# With name
geelark-cli browser automation tiktok-search --eid "557536075321468390" --schedule-at 1741846843 --search-keyword "hello" --name "myTask"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
