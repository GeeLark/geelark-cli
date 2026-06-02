# browser automation instagram-search

Create an Instagram automation task to search hashtags and browse posts.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--search-keywords <csv>` | Comma-separated search keywords (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Search Instagram hashtags
geelark-cli browser automation instagram-search --eid "557536075321468390" --schedule-at 1741846843 --search-keywords "hello,world"

# With name
geelark-cli browser automation instagram-search --eid "557536075321468390" --schedule-at 1741846843 --search-keywords "hello" --name "myTask"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
