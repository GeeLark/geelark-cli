# browser get-cookie

Query the cookies of a browser environment. Only one environment can be queried at a time.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Browser ID (required) |

## Examples

```bash
# Get browser cookies
geelark-cli browser get-cookie --id "browser_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `cookies` | string | Cookie data (JSON string) |
