# browser transfer

Transfer browser environments to a target account. Max 200 per request.

## Key Flags

| Flag | Description |
|------|-------------|
| `--env-ids <csv>` | Comma-separated browser IDs, max 200 (required) |
| `--username <text>` | Target user account email (required) |
| `--transfer-option <csv>` | Comma-separated options: name, proxy, tag, remark, files |

## Examples

```bash
# Transfer browsers
geelark-cli browser transfer --env-ids "id1,id2" --username "user@geelark.com"

# Transfer with options
geelark-cli browser transfer --env-ids "id1" --username "user@geelark.com" --transfer-option "name,proxy,tag,remark"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `successCount` | integer | Number of successful transfers |
| `failCount` | integer | Number of failed transfers |
| `failEnvIds` | array[string] | IDs of environments where transfers failed |

## Error Codes

| Code | Description |
|------|-------------|
| 40013 | Target user account does not exist |
| 43022 | Cannot transfer to own account |
| 43027 | Browser cannot be transferred |
| 40011 | Current user is not a paying user |
| 43028 | Sub-user lacks environment group permissions |
