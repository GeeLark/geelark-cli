# phone webhook get

Get the currently set webhook callback URL.

## Examples

```bash
geelark-cli phone webhook get
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `url` | string | The set callback URL |

## Error Codes

| Code | Description |
|------|-------------|
| 51001 | Callback URL not set |
