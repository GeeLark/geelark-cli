# phone create

Batch create cloud phones. Basic plan supports 1 at a time; Pro plan supports batch (max 100).

## Key Flags

| Flag | Description |
|------|-------------|
| `--mobile-type <text>` | Android version: Android 9/10/11/12/13/14/15/16 (default "Android 12") |
| `--charge-mode <n>` | Billing mode: 0=on-demand, 1=monthly |
| `--region <text>` | Region: cn, sgp, us (required) |
| `--data <json>` | JSON array of environment parameters (required) |

### Environment Parameters (in --data JSON)

| Parameter | Required | Description |
|-----------|----------|-------------|
| `profileName` | Yes | Cloud phone name |
| `proxyInformation` | No | Proxy string, e.g. socks5://user:pass@host:port |
| `refreshUrl` | No | Proxy refresh URL |
| `proxyQueryChannel` | No | Proxy detection: 1=ip-api, 2=IP2Location (default 2) |
| `proxyNumber` | No | Serial number of an added proxy |
| `dynamicProxy` | No | Saved dynamic proxy name |
| `dynamicProxyLocation` | No | Dynamic proxy country code |
| `mobileRegion` | No | Cloud phone region (follows proxy if empty) |
| `mobileProvince` | No | State (US only, follows proxy if empty) |
| `mobileCity` | No | City (US only, follows proxy if empty) |
| `mobileLanguage` | No | Language (default=English, baseOnIP=follows proxy) |
| `profileGroup` | No | Group name (auto-created if not exists) |
| `profileTags` | No | Tag names array (auto-created if not exists) |
| `profileNote` | No | Remark |
| `surfaceBrandName` | No | Phone brand (from brand-list) |
| `surfaceModelName` | No | Phone model (from brand-list) |
| `isTeamBrand` | No | Whether brand is team-uploaded |
| `netType` | No | Network: 0=Wi-Fi, 1=Mobile (Android 12/13/15 only) |
| `phoneNumber` | No | Custom phone number (auto-generated if empty) |
| `phoneName` | No | Device name (auto-generated if empty, not on Android 9/11) |

## Examples

```bash
# Create with proxy
geelark-cli phone create --region "sgp" --mobile-type "Android 12" --data "[{\"profileName\":\"myPhone\",\"proxyInformation\":\"socks5://user:pass@1.2.3.4:1080\"}]"

# Create with group and tags
geelark-cli phone create --region "cn" --mobile-type "Android 13" --data "[{\"profileName\":\"phone1\",\"profileGroup\":\"marketing\",\"profileTags\":[\"tiktok\"]}]"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total creation count |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `details[]` | array | Creation details |
| `details[].index` | integer | Creation index |
| `details[].code` | integer | Result code (0=success) |
| `details[].msg` | string | Result message |
| `details[].id` | string | Cloud phone ID |
| `details[].profileName` | string | Cloud phone name |
| `details[].envSerialNo` | string | Serial number |
| `details[].equipmentInfo` | object | Device info (same structure as list) |

## Error Codes

| Code | Description |
|------|-------------|
| 44001 | Batch creation requires Pro plan |
| 44002 | Plan environment limit reached |
| 44004 | Daily creation limit reached |
| 43029 | Selected model under maintenance |
| 45005 | Incorrect timezone setting |
