# proxy check

Check/detect a proxy to verify connectivity and get outbound IP information.

## Key Flags

| Flag | Description |
|------|-------------|
| `--type <text>` | Proxy type: socks5 (default), http, https |
| `--server <text>` | Proxy server address (required) |
| `--port <n>` | Proxy port (required) |
| `--username <text>` | Proxy username (optional) |
| `--password <text>` | Proxy password (optional) |
| `--channel <text>` | IP lookup source: IP-API or IP2Location (default IP2Location) |

## Examples

```bash
# Check a proxy
geelark-cli proxy check --type socks5 --server 185.162.130.86 --port 10000

# Check with auth and specific channel
geelark-cli proxy check --type socks5 --server 185.162.130.86 --port 10000 --username user --password pass --channel IP-API
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `detectStatus` | boolean | Whether the detection was successful |
| `message` | string | Reason for failure (if any) |
| `outboundIP` | string | Outbound IP address |
| `countryCode` | string | Country code of the outbound IP |
| `countryName` | string | Country name of the outbound IP |
| `subdivision` | string | State/Province of the outbound IP |
| `city` | string | City of the outbound IP |
| `timezone` | string | Timezone of the outbound IP |
| `isp` | string | ISP of the outbound IP |
