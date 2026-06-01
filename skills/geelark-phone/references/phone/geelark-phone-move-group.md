# phone move-group

Move cloud phones to a specified group. Max 100 phones. Use group-id "0" for ungrouped.

## Key Flags

| Flag | Description |
|------|-------------|
| `--env-ids <csv>` | Cloud phone IDs, max 100 (required) |
| `--group-id <text>` | Target group ID, "0" for ungrouped (required) |

## Examples

```bash
# Move to group
geelark-cli phone move-group --env-ids "id1,id2" --group-id "group_id"

# Move to ungrouped
geelark-cli phone move-group --env-ids "id1" --group-id "0"
```
