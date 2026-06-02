# browser api-status

Check whether the local Browser API is available.

## Examples

```bash
geelark-cli browser api-status
```

## Response Fields

> The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

On success, `code` is `0` and `msg` is `"success"`. No additional `data` fields.
