# proxy simple-add

Quick add a single proxy using flat flags. For batch creation, use `add`.

## Key Flags

| Flag | Description |
|------|-------------|
| `--scheme <text>` | Proxy type: socks5 (default), http, https |
| `--server <text>` | Proxy server address (required) |
| `--port <n>` | Proxy port (required) |
| `--username <text>` | Proxy username (optional) |
| `--password <text>` | Proxy password (optional) |
| `--proxy-query-channel <n>` | Detection channel: 1=IP-API, 2=IP2Location (default 2) |

## Examples

```bash
# Add without auth
geelark-cli proxy simple-add --scheme socks5 --server 192.3.8.1 --port 8000

# Add with auth
geelark-cli proxy simple-add --scheme socks5 --server 192.3.8.1 --port 8000 --username admin --password admin

# Add with specific detection channel
geelark-cli proxy simple-add --scheme http --server 1.2.3.4 --port 8080 --proxy-query-channel 1
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format). Response structure is identical to [`add`](geelark-proxy-add.md).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count (always 1 for simple-add) |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `successDetails[]` | array | Successful details |
| `successDetails[].index` | integer | Index of the provided proxy item |
| `successDetails[].id` | string | Proxy ID |
| `failDetails[]` | array | Failed details |
| `failDetails[].index` | integer | Index of the provided proxy item |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 45003 | Proxy not allowed |
| 45004 | Check proxy failed |
| 45007 | Proxy already exists |
