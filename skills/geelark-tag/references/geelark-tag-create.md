# tag create

Batch create tags. Each tag requires a name (required) and an optional color. Default color is white.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON array; each element contains `name` (required) and `color` (optional: white, red, blue, green, yellow, purple) |

## Examples

```bash
# Batch create
geelark-cli tag create --data "[{\"name\":\"marketing\",\"color\":\"red\"},{\"name\":\"sales\",\"color\":\"blue\"}]"

# Create with default color (white)
geelark-cli tag create --data "[{\"name\":\"defaultTag\"}]"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `successDetails[]` | array | Successful details |
| `successDetails[].id` | string | Tag ID |
| `successDetails[].name` | string | Tag name |
| `successDetails[].color` | string | Tag color |
| `failDetails[]` | array | Failed details |
| `failDetails[].code` | integer | Error code |
| `failDetails[].name` | string | Tag name |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 43020 | Tag name is empty |
| 43021 | Tag name already exists |
| 43023 | Tag color not supported |
