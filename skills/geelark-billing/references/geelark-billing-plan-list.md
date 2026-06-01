# billing plan-list

Get all available subscription plans with pricing and limits.

## Examples

```bash
geelark-cli billing plan-list
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format). Note: `data` is an array directly (not paginated).

| Field | Type | Description |
|-------|------|-------------|
| `[]` | array | Plan list |
| `[].id` | string | Profiles ID (used in plan-upgrade) |
| `[].price` | float | Price for one month |
| `[].level` | integer | Plan level: 0=Base, 1=Pro |
| `[].envNum` | integer | Max environment number |
| `[].freeTime` | integer | Free use minutes |
| `[].openEnvNumOneDay` | integer | Max environments opened per day |
| `[].createEnvNumOneDay` | integer | Max new environments created per day |
