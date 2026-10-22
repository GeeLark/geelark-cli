# phone start

Batch start cloud phones. Max 200 IDs per request.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs, max 200 (required) |
| `--width <n>` | Display width in px, 200-600 (default 336) |
| `--center <n>` | Center display: 0=no, 1=yes (default 1) |
| `--energy-saving <n>` | Energy-saving: 0=disabled, 1=enabled (auto shutdown after 30min idle). Deprecated: prefer `--energy-saving-after-minutes` |
| `--energy-saving-after-minutes <n>` | Enable energy-saving after N minutes idle (auto shutdown): 30, 60, 90; 0=disabled. Only sent when explicitly set |
| `--cleanup-sdcard-after-days <n>` | Auto-clean SD card data older than N days (speeds up startup): 3, 7; 0=disabled; omitted keeps team default. Only sent when explicitly set |
| `--dark-mode` | Dark mode (bool). Enables dark theme on cloud phone display |
| `--lang <text>` | Language, optional (supported: zh-CN, en-US, vi-VN, pt-BR, ru-RU, es-LA, uk-UA, ja-JP, zh-TW, id-ID, tr-TR, fr-FR). If omitted, the server-side per-domain default applies |
| `--material-tag-ids <csv>` | Material tag IDs, max 10 (requires OEM) |

## Examples

```bash
# Start phones
geelark-cli phone start --ids "id1,id2,id3"

# Start with display options
geelark-cli phone start --ids "id1" --width 480 --center 1 --energy-saving 1

# Start with a specific language
geelark-cli phone start --ids "id1" --lang "zh-CN"

# Start with energy-saving (auto shutdown after 60min idle) and SD card cleanup
geelark-cli phone start --ids "id1" --energy-saving-after-minutes 60 --cleanup-sdcard-after-days 7
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total requested count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `successDetails[]` | array | Successful details |
| `successDetails[].id` | string | Cloud phone ID |
| `successDetails[].url` | string | Remote URL for browser access |
| `successDetails[].chargingMethod` | string | Billing type (Per-minute usage, Monthly rental, Parallels) |
| `failDetails[]` | array | Failed details |
| `failDetails[].code` | integer | Error code |
| `failDetails[].id` | string | Cloud phone ID |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 43004 | Cloud phone has expired |
| 43007 | Cloud phone already in use |
| 43020 | Cloud phone currently unavailable |
| 43029 | Selected model under maintenance |
| 45002 | Cloud phone proxy unavailable |
| 47002 | Cloud phone resources insufficient |
| 47004 | Device does not exist |
