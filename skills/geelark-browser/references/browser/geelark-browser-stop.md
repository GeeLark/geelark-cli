# browser stop

Close a running browser environment by ID.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Browser ID (required) |

## Examples

```bash
# Stop a browser
geelark-cli browser stop --id "browser_id"
```

## Response Fields

> The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

On success, `code` is `0` and `msg` is `"success"`. No additional `data` fields.

## Error Codes

| Code | Description |
|------|-------------|
| -1 | Shutdown failed |
| 43028 | Sub-user lacks environment group permissions |
