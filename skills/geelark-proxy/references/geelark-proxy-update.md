# proxy update

Batch update proxies. Up to 100 items per request.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON array; each element contains `id` (required), `scheme` (required), `server` (required), `port` (required), `username` (optional), `password` (optional), `proxyQueryChannel` (optional, defaults to original) |

## Examples

```bash
# Batch update
geelark-cli proxy update --data "[{\"id\":\"proxy_id\",\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000,\"username\":\"admin\",\"password\":\"admin\",\"proxyQueryChannel\":2}]"

# Update multiple
geelark-cli proxy update --data "[{\"id\":\"id1\",\"scheme\":\"http\",\"server\":\"1.2.3.4\",\"port\":8080,\"proxyQueryChannel\":1}]"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count (duplicate IDs not counted) |
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
