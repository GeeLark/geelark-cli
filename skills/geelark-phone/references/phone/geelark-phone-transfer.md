# phone transfer

Transfer cloud phones to another account. Max 200 phones per request.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs, max 200 (required) |
| `--account <text>` | Target account email (required) |
| `--transfer-option <csv>` | Transfer options: name, proxy, tag, remark, files |

## Examples

```bash
# Transfer with all options
geelark-cli phone transfer --ids "id1,id2" --account "user@geelark.com" --transfer-option "name,proxy,tag,remark"

# Transfer without options
geelark-cli phone transfer --ids "id1" --account "user@geelark.com"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `[]` | array | Result array |
| `[].successCount` | integer | Successful count |
| `[].failCount` | integer | Failed count |
| `[].failEnvIds` | array[string] | Failed cloud phone IDs (in use or not found) |

## Error Codes

| Code | Description |
|------|-------------|
| 40013 | Target account not found |
| 43022 | Cannot transfer to self |
