# browser automation facebook-homepage

Create a Facebook automation task to create a homepage with specified name and category.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--main-page-name <text>` | Homepage name (required) |
| `--category <csv>` | Comma-separated homepage categories (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Create Facebook homepage
geelark-cli browser automation facebook-homepage --eid "557536075321468390" --schedule-at 1741846843 --main-page-name "myPage" --category "dev"

# With multiple categories
geelark-cli browser automation facebook-homepage --eid "557536075321468390" --schedule-at 1741846843 --main-page-name "myPage" --category "dev,business"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
