# tag simple-create

Quick create a single tag using flat flags. For batch creation, use `create`.

## Key Flags

| Flag | Description |
|------|-------------|
| `--name <text>` | Tag name, max 30 characters (required) |
| `--color <text>` | Tag color: white (default), red, blue, green, yellow, purple |

## Examples

```bash
# Name only (default color: white)
geelark-cli tag simple-create --name "marketing"

# Name and color
geelark-cli tag simple-create --name "marketing" --color red
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format). Response structure is identical to [`create`](geelark-tag-create.md).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count (always 1 for simple-create) |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `successDetails[]` | array | Successful details |
| `successDetails[].id` | string | Tag ID |
| `successDetails[].name` | string | Tag name |
| `successDetails[].color` | string | Tag color |
| `failDetails[]` | array | Failed details |
| `failDetails[].code` | integer | Error code |
| `failDetails[].name` | string | Tag name |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 43020 | Tag name is empty |
| 43021 | Tag name already exists |
| 43023 | Tag color not supported |
