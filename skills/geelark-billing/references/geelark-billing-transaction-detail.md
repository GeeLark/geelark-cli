# billing transaction-detail

Query billing transaction details. Only supports searching data within the last 3 days.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (optional, returns all if not specified) |
| `--start-at <n>` | Filter start time, second-level timestamp (last 3 days only) |
| `--end-at <n>` | Filter end time, second-level timestamp (last 3 days only) |
| `--limit <n>` | Number of records to return (default 100, max 1000) |
| `--last-flow-id <text>` | Last flow ID from previous request for pagination |

## Examples

```bash
# Recent transactions
geelark-cli billing transaction-detail --limit 20

# Filter by cloud phone
geelark-cli billing transaction-detail --id "phone_id" --limit 20

# Filter by time range
geelark-cli billing transaction-detail --start-at 1774593838 --end-at 1774593840

# Paginate with last flow ID
geelark-cli billing transaction-detail --last-flow-id "flow_id" --limit 100
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Page number |
| `pageSize` | integer | Page size |
| `list[]` | array | Transaction list |
| `list[].id` | string | Flow ID |
| `list[].envId` | string | Cloud phone ID |
| `list[].type` | integer | Usage type: 1=cloud phone, 2=RPA |
| `list[].chargeType` | integer | Billing type: 1=Points, 2=Balance, 3=Bonus, 4=Time add-on, 5=Bonus minutes, 6=Monthly rental, 7=Parallels, 8=Daily cap reached |
| `list[].amount` | float | Amount |
| `list[].usedTime` | integer | Usage duration in minutes |
| `list[].createdTime` | integer | Flow created time, second-level timestamp |
| `lastFlowId` | string | Last sequential ID for next page pagination |
