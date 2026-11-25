# phone set-gps

Set GPS location for cloud phones (batch). Longitude: [-180.0, 180.0], Latitude: [-90.0, 90.0]. Not supported on Android 16.
Optional extended fields per phone (all except latitude/longitude are optional; unset fields keep the phone's existing values): `bearing` (degrees, (0, 360]), `speed` (m/s, (0, 343]), `altitude` (meters, [-50000, 100000]), `accuracy` (meters, (0, 100]), `satellites` (0-36), `randomGps` object.
`randomGps`: `{"status":true,"radius":500,"unit":"m"}` — when enabled, each GPS push generates a random position within the radius around the given point; `radius` range is [1, 100] for `km` or [1, 100000] for `m`. The saved value is still the original point; the random range does not accumulate.
Extended field support varies by Android version (see table below); latitude/longitude and random GPS are supported on all models.

| Android | bearing | speed | altitude | accuracy | satellites |
|---------|---------|-------|----------|----------|-------------|
| 9 | Yes | Yes | Yes | No | No |
| 10 | Yes | No | Yes | Yes | No |
| 11 | No | No | Yes | Yes | No |
| 12 | Yes | Yes | Yes | Yes | Yes |
| 13 | Yes | Yes | Yes | Yes | Yes |

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON array of GPS data (required): `[{"id":"...","latitude":0.0,"longitude":0.0,"bearing":0.0,"speed":0.0,"altitude":0.0,"accuracy":0.0,"satellites":0,"randomGps":{"status":true,"radius":500,"unit":"m"}}]` |

## Examples

```bash
# Set GPS for one phone
geelark-cli phone set-gps --data "[{\"id\":\"phone_id\",\"latitude\":1.302,\"longitude\":103.875}]"

# Set GPS for multiple phones
geelark-cli phone set-gps --data "[{\"id\":\"id1\",\"latitude\":1.302,\"longitude\":103.875},{\"id\":\"id2\",\"latitude\":11.302,\"longitude\":104.875}]"

# Set GPS with extended fields and random GPS
geelark-cli phone set-gps --data "[{\"id\":\"phone_id\",\"latitude\":1.302,\"longitude\":103.875,\"bearing\":90,\"speed\":1.5,\"altitude\":15,\"accuracy\":10,\"satellites\":8,\"randomGps\":{\"status\":true,\"radius\":500,\"unit\":\"m\"}}]"
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
