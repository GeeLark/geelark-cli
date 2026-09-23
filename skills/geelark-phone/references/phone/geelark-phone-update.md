# phone update

Update cloud phone information. **Do not call while starting a cloud phone.**

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--name <text>` | New name, max 100 chars |
| `--remark <text>` | New remark, max 1500 chars |
| `--group-id <text>` | New group ID |
| `--tag-ids <csv>` | New tag IDs |
| `--proxy-id <text>` | Proxy ID |
| `--data <json>` | Additional JSON data (e.g. proxyConfig, chargeMode, phoneNumber) |

### Proxy Config (proxyConfig in --data)

Static proxy (typeId 1=socks5, 2=http, 3=https, -1=no proxy):

| Parameter | Required | Description |
|-----------|----------|-------------|
| `typeId` | Yes | Proxy type ID |
| `server` | Yes | Proxy server hostname |
| `port` | Yes | Proxy server port |
| `username` | Yes | Proxy server username |
| `password` | Yes | Proxy server password |

Dynamic proxy (typeId 21=IPHTML, 22=kookeey, 23=Lumatuo(BrightData), 24=Proxyma, 25=SmartProxy, 26=RolaIP, 27=NodeMaven, 29=KookeeyMobile, 30=GProxy Residential). Set `useProxyCfg: true` to reuse the already configured proxy without host/port:

| Parameter | Required | Description |
|-----------|----------|-------------|
| `useProxyCfg` | Yes | Use the already configured proxy |
| `typeId` | Yes | Proxy type ID (dynamic list above) |
| `protocol` | No | Proxy protocol: 1=SOCKS5, 2=HTTP |
| `server`/`port`/`username`/`password` | No | Proxy connection details |
| `country` | No | Dynamic proxy country code |
| `region` | No | Dynamic proxy region |
| `city` | No | Dynamic proxy city |
| `sessionType` | No | GProxy session type: 1=Sticky, 2=Rotating (default 1) |
| `filterMode` | No | GProxy filter mode: 0=MaxPoolSize, 1=Quality, 2=QualitySpeed, 3=Speed |

## Examples

```bash
# Update name and remark
geelark-cli phone update --id "phone_id" --name "new name" --remark "new remark"

# Update group and tags
geelark-cli phone update --id "phone_id" --group-id "group_id" --tag-ids "tag1,tag2"

# Update proxy via proxyConfig
geelark-cli phone update --id "phone_id" --data "{\"proxyConfig\":{\"typeId\":1,\"server\":\"1.2.3.4\",\"port\":1080,\"username\":\"u\",\"password\":\"p\"}}"

# Update charge mode
geelark-cli phone update --id "phone_id" --data "{\"chargeMode\":1}"

# Update proxy to GProxy residential dynamic proxy
geelark-cli phone update --id "phone_id" --data "{\"proxyConfig\":{\"useProxyCfg\":true,\"typeId\":30,\"sessionType\":1,\"filterMode\":1}}"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `failDetails[]` | array | Tag addition failure info (only present if some tags fail) |
| `failDetails[].code` | integer | Error code |
| `failDetails[].id` | string | Tag ID |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 43022 | Tag does not exist |
| 43032 | Group does not exist |
| 45003 | Proxy region not allowed |
| 45004 | Proxy check failed |
| 45008 | Proxy type not allowed |
