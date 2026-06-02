# browser check-status

Check whether a specific browser is running. Returns status (open/close) and debug port.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Browser ID (required) |

## Examples

```bash
# Check browser status
geelark-cli browser check-status --id "browser_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | string | Browser status: `"open"` or `"close"` |
| `debugPort` | integer | Debug port (available when open) |
