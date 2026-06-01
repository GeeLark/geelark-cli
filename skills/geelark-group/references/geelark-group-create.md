# group create

Batch create groups. Each group requires a name (required) and an optional remark.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON array; each element contains `name` (required) and `remark` (optional) |

## Examples

```bash
# Batch create
geelark-cli group create --data "[{\"name\":\"marketing\",\"remark\":\"Marketing team\"},{\"name\":\"sales\"}]"

# Create a single group (simple-create is recommended for this)
geelark-cli group create --data "[{\"name\":\"marketing\"}]"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `successDetails[]` | array | Successful details |
| `successDetails[].id` | string | Newly created group ID |
| `successDetails[].name` | string | Group name |
| `failDetails[]` | array | Failed details |
| `failDetails[].code` | integer | Error code |
| `failDetails[].name` | string | Failed group name |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 43030 | Group name is empty |
| 43031 | Group name already exists |
