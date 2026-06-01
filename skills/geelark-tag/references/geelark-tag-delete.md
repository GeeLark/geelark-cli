# tag delete

Delete tags by ID. Supports batch deletion.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Tag IDs to delete, comma-separated (required) |

## Examples

```bash
# Delete a single tag
geelark-cli tag delete --ids "tag_id"

# Batch delete
geelark-cli tag delete --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `failDetails[]` | array | Failed details |
| `failDetails[].code` | integer | Error code |
| `failDetails[].id` | string | Tag ID |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 43022 | Tag does not exist |
