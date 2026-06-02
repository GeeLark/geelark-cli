# browser automation facebook-friends

Create a Facebook automation task to add recommended friends.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Add recommended friends
geelark-cli browser automation facebook-friends --eid "557536075321468390" --schedule-at 1741846843

# With name
geelark-cli browser automation facebook-friends --eid "557536075321468390" --schedule-at 1741846843 --name "myTask"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
