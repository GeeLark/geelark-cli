# tag simple-update

Quick update a single tag using flat flags. For batch updates, use `update`.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Tag ID (required) |
| `--name <text>` | New tag name, max 30 characters (optional) |
| `--color <text>` | New tag color: white, red, blue, green, yellow, purple (optional) |

At least one of `--name` or `--color` must be provided.

## Examples

```bash
# Update name
geelark-cli tag simple-update --id "tag_id" --name "sales"

# Update color
geelark-cli tag simple-update --id "tag_id" --color blue

# Update both
geelark-cli tag simple-update --id "tag_id" --name "sales" --color blue
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format). Response structure is identical to [`update`](geelark-tag-update.md).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count (always 1 for simple-update) |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |

## Error Codes

| Code | Description |
|------|-------------|
| 43020 | Tag name is empty |
| 43022 | Tag does not exist |
| 43023 | Tag color not supported |
