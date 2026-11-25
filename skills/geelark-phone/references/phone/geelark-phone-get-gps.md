# phone get-gps

Get GPS information of cloud phones, including latitude/longitude and extended params (bearing, speed, altitude, accuracy, satellites) plus the random GPS setting.
Extended params are the currently effective values on the phone: when random GPS is enabled, the returned latitude/longitude is the randomized actual position, not the original point set. Unsupported or unset extended params usually return 0. `randomGps` is the setting saved on the platform; it is not returned if never set.

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
| `list[].latitude` | float | Latitude (randomized actual position when random GPS is enabled) |
| `list[].longitude` | float | Longitude (randomized actual position when random GPS is enabled) |
| `list[].bearing` | float | Bearing in degrees, currently effective value |
| `list[].speed` | float | Speed in m/s, currently effective value |
| `list[].altitude` | float | Altitude in meters, currently effective value |
| `list[].accuracy` | float | Location accuracy in meters, currently effective value |
| `list[].satellites` | integer | Number of satellites, currently effective value |
| `list[].randomGps.status` | boolean | Whether random GPS is enabled (absent if never set) |
| `list[].randomGps.radius` | integer | Random radius |
| `list[].randomGps.unit` | string | Radius unit: `km` or `m` |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
