# phone clone

Clone a cloud phone. Retains country, timezone, language, and GPS. Applications and data will be cleared.

## Key Flags

| Flag | Description |
|------|-------------|
| `--env-id <text>` | Cloud phone ID to clone (required) |
| `--amount <n>` | Number of clones, 1-100 (default 1) |
| `--group-id <text>` | Target group ID (ungrouped if not specified) |
| `--clone-name` | Clone the name |
| `--clone-remark` | Clone the remark |
| `--clone-tag` | Clone the tags |
| `--clone-proxy` | Clone the proxy |
| `--clone-net-type` | Clone the network type |

## Examples

```bash
# Clone 1 phone
geelark-cli phone clone --env-id "phone_id" --amount 1

# Clone with options
geelark-cli phone clone --env-id "phone_id" --amount 3 --group-id "group_id" --clone-name --clone-proxy --clone-tag
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `ids` | array[string] | Cloned cloud phone IDs |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 43032 | Group does not exist |
| 44001 | Pro plan required |
| 44002 | Plan environment limit reached |
| 44004 | Daily creation limit reached |
| 44006 | Insufficient cloud phone inventory |
| 43038 | Device model deleted |
