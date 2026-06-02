# browser automation cookie-bot

Create a Cookie Bot automation task that visits specified webpages to collect cookies.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--pages <csv>` | Comma-separated webpage URLs to visit (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Create Cookie Bot task
geelark-cli browser automation cookie-bot --eid "557536075321468390" --schedule-at 1741846843 --pages "https://a.com,https://b.com"

# With name
geelark-cli browser automation cookie-bot --eid "557536075321468390" --schedule-at 1741846843 --pages "https://a.com" --name "myTask"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
