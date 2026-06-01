# phone delete

Batch delete cloud phones. Phones must be stopped first. Max 100 IDs per request.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs, max 100 (required) |

## Examples

```bash
geelark-cli phone delete --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total requested count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `failDetails[]` | array | Failed details |
| `failDetails[].code` | integer | Error code |
| `failDetails[].id` | string | Cloud phone ID |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 43009 | Cloud phone is started (stop first) |
| 43010 | Cloud phone is starting (wait) |
| 43021 | Cloud phone is in use, try again later |
