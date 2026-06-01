# phone list

List cloud phones with optional filters. Supports pagination and multiple filter criteria.

## Key Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |
| `--ids <csv>` | Filter by cloud phone IDs (max 100, ignores page/pageSize) |
| `--serial-name <text>` | Filter by cloud phone name |
| `--remark <text>` | Filter by remark |
| `--group-name <text>` | Filter by group name |
| `--tags <csv>` | Filter by tag names |
| `--charge-mode <n>` | Charge mode: 0=pay-per-minute, 1=monthly |
| `--open-status <n>` | Power state: 0=off, 1=on |
| `--proxy-ids <csv>` | Filter by proxy IDs (max 10) |
| `--serial-nos <csv>` | Filter by serial numbers (max 100) |

## Examples

```bash
# List all phones
geelark-cli phone list --page 1 --page-size 10

# Filter by name
geelark-cli phone list --serial-name "myPhone"

# Filter by status
geelark-cli phone list --open-status 1 --charge-mode 0

# Filter by IDs
geelark-cli phone list --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `items[]` | array | Cloud phone list |
| `items[].id` | string | Cloud phone ID |
| `items[].serialName` | string | Cloud phone name |
| `items[].serialNo` | string | Serial number |
| `items[].remark` | string | Remark |
| `items[].status` | integer | Status: 0=started, 1=starting, 2=shut down |
| `items[].chargeMode` | integer | Charge mode: 0=pay-per-minute, 1=monthly |
| `items[].hasBind` | bool | Whether bound to monthly subscription |
| `items[].monthlyExpire` | integer | Monthly expiration timestamp (seconds) |
| `items[].rpaStatus` | integer | RPA running: 1=running, 0=not running |
| `items[].createTime` | integer | Creation timestamp (seconds) |
| `items[].group` | object | Group info |
| `items[].group.id` | string | Group ID |
| `items[].group.name` | string | Group name |
| `items[].group.remark` | string | Group remark |
| `items[].tags[]` | array | Tag list |
| `items[].tags[].name` | string | Tag name |
| `items[].equipmentInfo` | object | Device info |
| `items[].equipmentInfo.countryName` | string | Country |
| `items[].equipmentInfo.phoneNumber` | string | Phone number |
| `items[].equipmentInfo.enableSim` | integer | SIM enabled: 0=no, 1=yes |
| `items[].equipmentInfo.imei` | string | IMEI |
| `items[].equipmentInfo.osVersion` | string | OS version |
| `items[].equipmentInfo.wifiBssid` | string | Wi-Fi MAC |
| `items[].equipmentInfo.mac` | string | Phone Wi-Fi MAC |
| `items[].equipmentInfo.bluetoothMac` | string | Bluetooth MAC |
| `items[].equipmentInfo.timeZone` | string | Timezone |
| `items[].equipmentInfo.deviceBrand` | string | Brand |
| `items[].equipmentInfo.deviceModel` | string | Model |
| `items[].equipmentInfo.deviceName` | string | Device name |
| `items[].equipmentInfo.netType` | integer | Network: 0=Wi-Fi, 1=Mobile |
| `items[].equipmentInfo.language` | string | Language |
| `items[].equipmentInfo.province` | string | Province |
| `items[].equipmentInfo.city` | string | City |
| `items[].proxy` | object | Proxy info |
| `items[].proxy.type` | string | Proxy type (socks5, http, https) |
| `items[].proxy.server` | string | Proxy server |
| `items[].proxy.port` | integer | Proxy port |
| `items[].proxy.username` | string | Proxy username |
| `items[].proxy.password` | string | Proxy password |
