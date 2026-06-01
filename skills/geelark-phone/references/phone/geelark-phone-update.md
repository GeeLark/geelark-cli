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
