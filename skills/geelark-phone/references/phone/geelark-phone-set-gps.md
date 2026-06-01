# phone set-gps

Set GPS location for cloud phones (batch). Longitude: [-180.0, 180.0], Latitude: [-90.0, 90.0]. Not supported on Android 16.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON array of GPS data (required): `[{"id":"...","latitude":0.0,"longitude":0.0}]` |

## Examples

```bash
# Set GPS for one phone
geelark-cli phone set-gps --data "[{\"id\":\"phone_id\",\"latitude\":1.302,\"longitude\":103.875}]"

# Set GPS for multiple phones
geelark-cli phone set-gps --data "[{\"id\":\"id1\",\"latitude\":1.302,\"longitude\":103.875},{\"id\":\"id2\",\"latitude\":11.302,\"longitude\":104.875}]"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total requested count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 43012 | Latitude/longitude range error |
