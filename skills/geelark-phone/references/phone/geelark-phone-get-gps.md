# phone get-gps

Get GPS information (latitude and longitude) of cloud phones.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs (required) |

## Examples

```bash
geelark-cli phone get-gps --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total requested count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `list[]` | array | GPS info list |
| `list[].id` | string | Cloud phone ID |
| `list[].latitude` | float | Latitude |
| `list[].longitude` | float | Longitude |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
