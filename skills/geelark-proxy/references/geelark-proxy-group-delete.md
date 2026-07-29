# proxy group-delete

Delete proxy groups in batch.

After a proxy group is deleted, proxies in that group are automatically moved to the ungrouped category. Deleting an already-deleted group, a nonexistent group, or group ID `"0"` is treated as successful.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Proxy group IDs to delete, comma-separated (required). Duplicate IDs are counted only once |

## Examples

```bash
# Delete a single group
geelark-cli proxy group-delete --ids "123456789012345678"

# Batch delete
geelark-cli proxy group-delete --ids "123456789012345678,223456789012345678"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total processed (duplicate IDs counted only once) |
| `successAmount` | integer | Successfully deleted count |
| `failAmount` | integer | Failed count |
| `failDetails[]` | array | Failed details |
| `failDetails[].id` | string | Proxy group ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

## Notes

- Deleting group ID `"0"` (ungrouped) is treated as successful but does nothing
- Proxies in deleted groups are moved to the ungrouped category, not deleted
