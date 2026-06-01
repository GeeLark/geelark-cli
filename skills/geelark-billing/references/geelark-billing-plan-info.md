# billing plan-info

Get current subscription plan information. Rate limit: 1 request/minute.

## Examples

```bash
geelark-cli billing plan-info
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `plan` | integer | Subscription type: 0=Base, 1=Pro |
| `profiles` | integer | Total number of profiles |
| `monthlyRental` | integer | Number of monthly rental devices |
| `parallels` | integer | Max simultaneous open cloud phones |
| `expirationTime` | integer | Plan expiration timestamp (seconds) |
| `monthlyFee` | float | Monthly billing amount |
| `availableProfiles` | integer | Number of available profiles |
| `availableMonthlyRentals` | integer | Number of available monthly rental devices |
