# phone adb get-info

Get ADB connection information (IP, port, password) for cloud phones.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs, max 200 (required) |

## Examples

```bash
geelark-cli phone adb get-info --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `items[]` | array | Per-phone ADB info |
| `items[].id` | string | Cloud phone ID |
| `items[].code` | integer | Result code (0=success) |
| `items[].ip` | string | Connection IP |
| `items[].port` | string | Port |
| `items[].pwd` | string | ADB password |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 49001 | ADB is not enabled |
| 49002 | Device does not support ADB |
