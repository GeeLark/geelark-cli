# group simple-update

Quick update a single group using flat flags instead of JSON. For batch updates, use `update`.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Group ID (required) |
| `--name <text>` | New group name, max 50 characters (optional) |
| `--remark <text>` | New group remark, max 500 characters (optional) |

At least one of `--name` or `--remark` must be provided.

## Examples

```bash
# Update name
geelark-cli group simple-update --id "528995439832269824" --name "marketing-v2"

# Update name and remark
geelark-cli group simple-update --id "528995439832269824" --name "marketing-v2" --remark "Updated"

# Update remark only
geelark-cli group simple-update --id "528995439832269824" --remark "New remark"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format). Response structure is identical to [`update`](geelark-group-update.md).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count (always 1 for simple-update) |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |

## Error Codes

| Code | Description |
|------|-------------|
| 43030 | Group name is empty |
| 43032 | Group not found |
| 43035 | Cannot operate on "Ungrouped" |
