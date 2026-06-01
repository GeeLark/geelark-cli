# tag update

Batch update tag information (name, color). Each entry requires a tag ID (required) and the fields to modify.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON array; each element contains `id` (required), `name` (optional), `color` (optional) |

## Examples

```bash
# Batch update
geelark-cli tag update --data "[{\"id\":\"tag_id\",\"name\":\"newName\",\"color\":\"blue\"}]"

# Update name only
geelark-cli tag update --data "[{\"id\":\"tag_id\",\"name\":\"sales\"}]"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |

## Error Codes

| Code | Description |
|------|-------------|
| 43020 | Tag name is empty |
| 43022 | Tag does not exist |
| 43023 | Tag color not supported |
