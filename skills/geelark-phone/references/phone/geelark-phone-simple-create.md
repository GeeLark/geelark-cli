# phone simple-create

Quick create a single cloud phone using flat flags. For batch creation, use `create`.

## Key Flags

| Flag | Description |
|------|-------------|
| `--profile-name <text>` | Cloud phone name (required) |
| `--mobile-type <text>` | Android version (default "Android 12") |
| `--charge-mode <n>` | Billing mode: 0=on-demand, 1=monthly |
| `--region <text>` | Region: cn, sgp, us (required) |
| `--proxy-information <text>` | Proxy string, e.g. socks5://user:pass@host:port |
| `--proxy-query-channel <n>` | Proxy detection: 1=ip-api, 2=IP2Location |
| `--proxy-number <n>` | Serial number of an added proxy |
| `--profile-group <text>` | Group name (auto-created if not exists) |
| `--profile-tags <csv>` | Tag names (auto-created if not exists) |
| `--profile-note <text>` | Remark |
| `--net-type <n>` | Network: 0=Wi-Fi, 1=Mobile (Android 12/13/15 only) |
| `--phone-number <text>` | Custom phone number (auto-generated if empty) |
| `--phone-name <text>` | Device name (auto-generated if empty, not on Android 9/11) |

At least one of `--proxy-information` or `--proxy-number` is required.

## Examples

```bash
# Create with proxy info
geelark-cli phone simple-create --region sgp --mobile-type "Android 12" --profile-name "myPhone" --proxy-information "socks5://user:pass@1.2.3.4:1080"

# Create with proxy number and group
geelark-cli phone simple-create --region cn --profile-name "phone1" --proxy-number 1 --profile-group "marketing" --profile-tags "tiktok,sales"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format). Response structure is identical to [`create`](geelark-phone-create.md).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total creation count (always 1) |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `details[]` | array | Creation details |
| `details[].index` | integer | Creation index |
| `details[].code` | integer | Result code (0=success) |
| `details[].msg` | string | Result message |
| `details[].id` | string | Cloud phone ID |
| `details[].profileName` | string | Cloud phone name |
| `details[].envSerialNo` | string | Serial number |
| `details[].equipmentInfo` | object | Device info |

## Error Codes

| Code | Description |
|------|-------------|
| 44001 | Batch creation requires Pro plan |
| 44002 | Plan environment limit reached |
| 44004 | Daily creation limit reached |
| 43029 | Selected model under maintenance |
| 45005 | Incorrect timezone setting |
