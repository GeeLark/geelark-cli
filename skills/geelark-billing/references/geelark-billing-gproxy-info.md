# billing gproxy-info

Query the team's GProxy residential proxy traffic information. No flags required.
Used traffic is fetched in real time from the proxy provider; results within the last 3 minutes may be served from cache. When no residential proxy has been purchased, `residentialStatus` is `0` and both traffic values are `0`.

## Examples

```bash
geelark-cli billing gproxy-info
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `residentialStatus` | integer | Residential proxy status: `0` not purchased, `1` initializing, `2` available |
| `residentialUsedBytes` | integer | Used traffic, in bytes |
| `residentialAvailBytes` | integer | Available (remaining) traffic, in bytes |
