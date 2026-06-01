# tag list

List tags with optional filters by ID, name, or color.

## Key Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |
| `--ids <csv>` | Filter by tag IDs, comma-separated |
| `--names <csv>` | Filter by tag names, comma-separated |
| `--colors <csv>` | Filter by tag colors, comma-separated (white, red, blue, green, yellow, purple) |

## Examples

```bash
# List all tags
geelark-cli tag list --page 1 --page-size 50

# Filter by name
geelark-cli tag list --names "marketing,sales"

# Filter by color
geelark-cli tag list --colors "red,blue"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page number |
| `pageSize` | integer | Page size |
| `list[]` | array | Tag list |
| `list[].id` | string | Tag ID |
| `list[].name` | string | Tag name |
| `list[].color` | string | Tag color |
| `failDetails[]` | array | Failure details (on query) |
| `failDetails[].code` | integer | Error code |
| `failDetails[].type` | integer | Failure type: 1=Tag ID, 2=Tag name, 3=Tag color |
| `failDetails[].param` | string | Failed parameter |
| `failDetails[].msg` | string | Failure message |

## Error Codes

| Code | Description |
|------|-------------|
| 43022 | Tag does not exist |
| 43023 | Tag color does not exist |
| 43024 | Tag name does not exist |
