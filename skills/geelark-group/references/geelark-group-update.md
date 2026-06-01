# group update

Batch update group information (name, remark). Each entry requires a group ID (required) and the fields to modify.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON array; each element contains `id` (required), `name` (optional), `remark` (optional) |

## Examples

```bash
# Batch update
geelark-cli group update --data "[{\"id\":\"528995439832269824\",\"name\":\"newName\",\"remark\":\"newRemark\"}]"

# Update remark only
geelark-cli group update --data "[{\"id\":\"528995439832269824\",\"remark\":\"updated\"}]"
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
| 43030 | Group name is empty |
| 43032 | Group not found |
| 43035 | Cannot operate on "Ungrouped" |
