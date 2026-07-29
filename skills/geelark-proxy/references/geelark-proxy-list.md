# proxy list

List all proxies with pagination. Supports filtering by ID and by proxy group.

## Key Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |
| `--ids <csv>` | Filter by proxy IDs, comma-separated |
| `--proxy-group-id <text>` | Filter by proxy group ID (`"0"` = ungrouped) |

## Examples

```bash
# List proxies
geelark-cli proxy list --page 1 --page-size 10

# Filter by ID
geelark-cli proxy list --ids "493188072704313353"

# Filter by proxy group
geelark-cli proxy list --page 1 --page-size 10 --proxy-group-id "123456789012345678"

# Show only ungrouped proxies
geelark-cli proxy list --page 1 --page-size 10 --proxy-group-id "0"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page number |
| `pageSize` | integer | Page size |
| `list[]` | array | Proxy list |
| `list[].id` | string | Proxy ID |
| `list[].serialNo` | integer | Proxy serial number |
| `list[].scheme` | string | Proxy type: socks5, http, https |
| `list[].server` | string | Proxy address |
| `list[].port` | integer | Proxy port |
| `list[].username` | string | Proxy username |
| `list[].password` | string | Proxy password |
| `list[].proxyGroupId` | string | Proxy group ID (`"0"` = ungrouped) |
