# proxy simple-update

Quick update a single proxy using flat flags. For batch updates, use `update`.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Proxy ID (required) |
| `--scheme <text>` | Proxy type: socks5 (default), http, https |
| `--server <text>` | Proxy server address (required) |
| `--port <n>` | Proxy port (required) |
| `--username <text>` | Proxy username (optional) |
| `--password <text>` | Proxy password (optional) |
| `--proxy-query-channel <n>` | Detection channel: 1=IP-API, 2=IP2Location (default 2) |

## Examples

```bash
# Update server and port
geelark-cli proxy simple-update --id "proxy_id" --scheme socks5 --server 192.3.8.1 --port 8000

# Update with auth and detection channel
geelark-cli proxy simple-update --id "proxy_id" --scheme socks5 --server 192.3.8.1 --port 8000 --username admin --password admin --proxy-query-channel 1
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format). Response structure is identical to [`update`](geelark-proxy-update.md).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count (always 1 for simple-update) |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `failDetails[]` | array | Failed details |
| `failDetails[].id` | string | Proxy ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 40005 | Proxy not found |
| 45003 | Proxy not allowed |
| 45004 | Check proxy failed |
| 45007 | Proxy already exists |
| 45008 | Proxy type not allowed |
