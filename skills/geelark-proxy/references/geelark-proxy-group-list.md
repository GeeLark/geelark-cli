# proxy group-list

List proxy groups with optional fuzzy search by name.

The virtual "Ungrouped Proxies" category (`id="0"`) always appears on the first page and is included in `total`.

## Key Flags

| Flag | Description |
|------|-------------|
| `--name <text>` | Proxy group name (fuzzy search) |
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size (required, e.g. 20) |

## Examples

```bash
# List all groups
geelark-cli proxy group-list --page 1 --page-size 20

# Search by name
geelark-cli proxy group-list --page 1 --page-size 20 --name "Business"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count, including the virtual ungrouped category |
| `page` | integer | Current page number |
| `pageSize` | integer | Page size |
| `list[]` | array | Proxy group list |
| `list[].id` | string | Proxy group ID. `"0"` represents ungrouped proxies |
| `list[].name` | string | Proxy group name |
