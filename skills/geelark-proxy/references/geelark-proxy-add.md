# proxy add

Batch add proxies. Duplicate proxies will not be added. Up to 100 items per request.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON array; each element contains `scheme` (required), `server` (required), `port` (required), `username` (optional), `password` (optional), `proxyQueryChannel` (optional, 1=IP-API, 2=IP2Location), `proxyGroupId` (optional, group ID or `"0"` for ungrouped) |

## Examples

```bash
# Batch add
geelark-cli proxy add --data "[{\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000,\"username\":\"admin\",\"password\":\"admin\",\"proxyQueryChannel\":2}]"

# Add multiple proxies
geelark-cli proxy add --data "[{\"scheme\":\"http\",\"server\":\"1.2.3.4\",\"port\":8080,\"proxyQueryChannel\":1},{\"scheme\":\"socks5\",\"server\":\"5.6.7.8\",\"port\":1080,\"proxyQueryChannel\":2}]"

# Add proxies into a group
geelark-cli proxy add --data "[{\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000,\"proxyGroupId\":\"123456789012345678\"}]"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `successDetails[]` | array | Successful details |
| `successDetails[].index` | integer | Index of the provided proxy item |
| `successDetails[].id` | string | Proxy ID (same ID if proxy info is identical) |
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
