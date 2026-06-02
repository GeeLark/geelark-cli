# browser start

Launch a browser environment by ID. Optionally specify a webhook URL to receive a callback when the browser finishes starting.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Browser ID (required) |
| `--webhook <url>` | Callback URL to notify after browser startup completes |

## Webhook Callback

When a webhook URL is provided, a POST request will be sent after the browser startup task completes:

```json
{
  "event": "browser_start",
  "timestamp": 1776147008407,
  "data": {
    "id": "612342716134614477",
    "status": "success",
    "debugPort": 11019,
    "ipCheckPass": true
  }
}
```

## Examples

```bash
# Start a browser
geelark-cli browser start --id "browser_id"

# Start with webhook
geelark-cli browser start --id "browser_id" --webhook "http://localhost:3001"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `debugPort` | integer | Debug port for browser connection |

## Error Codes

| Code | Description |
|------|-------------|
| -1 | Startup failed |
| 43007 | Environment already in use |
| 43008 | Maximum open environments reached |
| 46003 | Environment not included in plan |
| 43028 | Sub-user lacks environment group permissions |
| 90002 | Environment does not exist |
| 90003 | Insufficient disk space |
