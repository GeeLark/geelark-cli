# browser automation youtube-watch

Create a YouTube automation task to search and watch videos, with optional title and comment.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--search-keyword <text>` | Search keyword (required) |
| `--title <text>` | Video title (required) |
| `--comment <text>` | Comment (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Create YouTube watch task
geelark-cli browser automation youtube-watch --eid "557536075321468390" --schedule-at 1741846843 --search-keyword "hello" --title "myTitle" --comment "myComment"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
