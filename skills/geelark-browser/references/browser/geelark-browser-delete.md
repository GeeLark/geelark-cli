# browser delete

Delete browser environments by IDs. Max 100 per request. Ensure browsers are closed first.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Comma-separated browser IDs (required) |

## Examples

```bash
# Delete browsers
geelark-cli browser delete --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `successIds` | array[string] | IDs of successfully deleted browsers |
| `busyIds` | array[string] | IDs of browsers currently in use |
| `serverErrIds` | array[string] | IDs of browsers with server errors |

## Error Codes

| Code | Description |
|------|-------------|
| 43028 | Sub-user lacks environment group permissions |
