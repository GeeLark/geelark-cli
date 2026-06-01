# proxy delete

Delete proxies by ID. Supports batch deletion. Up to 100 IDs per request.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Proxy IDs to delete, comma-separated (required) |

## Examples

```bash
# Delete a single proxy
geelark-cli proxy delete --ids "id1"

# Batch delete
geelark-cli proxy delete --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count (duplicate IDs not counted) |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `failDetails[]` | array | Failed details |
| `failDetails[].id` | string | Proxy ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

## Notes

- **Proxies bound to a cloud phone cannot be deleted**: Returns error `40010`

## Error Codes

| Code | Description |
|------|-------------|
| 40005 | Proxy not found |
| 40010 | Proxy is bound to an environment |
