# browser automation x-post

Create an X (Twitter) automation task to post a tweet with specified content.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--content <text>` | Tweet content (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Post a tweet
geelark-cli browser automation x-post --eid "557536075321468390" --schedule-at 1741846843 --content "hello"

# With name
geelark-cli browser automation x-post --eid "557536075321468390" --schedule-at 1741846843 --content "hello" --name "myTask"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
