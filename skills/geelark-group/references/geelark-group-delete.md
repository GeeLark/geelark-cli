# group delete

Delete groups by ID. Supports batch deletion.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Group IDs to delete, comma-separated (required) |

## Examples

```bash
# Delete a single group
geelark-cli group delete --ids "528995439832269824"

# Batch delete
geelark-cli group delete --ids "528995439832269824,528985080069096448"
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
| `failDetails[].id` | string | Failed group ID |
| `failDetails[].msg` | string | Error message |

## Notes

- **Deleting a group does not delete phones**: Phones in a deleted group are moved to "Ungrouped"
- "Ungrouped" cannot be deleted; attempting so returns `43035`

## Error Codes

| Code | Description |
|------|-------------|
| 43032 | Group not found |
| 43035 | Cannot operate on "Ungrouped" |
