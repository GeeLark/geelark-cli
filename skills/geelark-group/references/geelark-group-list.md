# group list

Query group list. Supports filtering by ID, name, or remark.

## Key Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |
| `--ids <csv>` | Filter by group IDs, comma-separated |
| `--names <csv>` | Filter by group names, comma-separated |

## Examples

```bash
# View all groups
geelark-cli group list --page 1 --page-size 50

# Query by ID
geelark-cli group list --ids "528995439832269824,528985080069096448"

# Query by name
geelark-cli group list --names "marketing,sales"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page number |
| `pageSize` | integer | Page size |
| `list[]` | array | Group list |
| `list[].id` | string | Group ID |
| `list[].name` | string | Group name |
| `list[].remark` | string | Group remark |

## Error Codes

| Code | Description |
|------|-------------|
| 43032 | Group ID not found |
| 43033 | Group name not found |
| 43034 | Group remark not found |
