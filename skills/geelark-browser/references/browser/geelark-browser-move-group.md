# browser move-group

Move browser environments to a specified group. Max 100 per request.

## Key Flags

| Flag | Description |
|------|-------------|
| `--env-ids <csv>` | Comma-separated browser IDs (required) |
| `--group-id <text>` | Target group ID, `"0"` for ungrouped (required) |

## Examples

```bash
# Move browsers to a group
geelark-cli browser move-group --env-ids "id1,id2" --group-id "group_id"

# Move to ungrouped
geelark-cli browser move-group --env-ids "id1,id2" --group-id "0"
```

## Response Fields

> The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

On success, `code` is `0` and `msg` is `"success"`. No additional `data` fields.
