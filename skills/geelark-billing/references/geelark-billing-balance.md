# billing balance

Query account balance. Rate limit: 10 requests/minute.

## Examples

```bash
geelark-cli billing balance
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `balance` | float | Account balance |
| `giftMoney` | float | Gifted amount |
| `availableTimeAddOn` | integer | Remaining time add-on minutes |
