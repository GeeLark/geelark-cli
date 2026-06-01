---
name: geelark-billing
version: 1.0.0
description: "GeeLark billing management: query balance, transaction details, plan info, upgrade/renew plans, and buy time add-ons. Use when managing account billing and subscriptions."
metadata:
  requires:
    bins: ["geelark-cli"]
  cliHelp: "geelark-cli billing --help"
---

# billing — Billing Management

**CRITICAL — MUST read [`../geelark-shared/SKILL.md`](../geelark-shared/SKILL.md) first for authentication and configuration handling.**

## Command Overview

```bash
geelark-cli billing <command> [flags]
```

| Command | Description |
|---------|-------------|
| `balance` | Query account balance |
| `transaction-detail` | Query billing transaction details |
| `plan-list` | Get all available plans |
| `plan-info` | Get current subscription plan info |
| `plan-upgrade` | Upgrade subscription plan |
| `plan-renew` | Renew subscription plan |
| `buy-time-addon` | Buy time add-on minutes |

## Which Command to Use

| Want to | Command |
|---------|---------|
| Check balance | [`balance`](references/geelark-billing-balance.md) |
| View transaction history | [`transaction-detail`](references/geelark-billing-transaction-detail.md) |
| See available plans | [`plan-list`](references/geelark-billing-plan-list.md) |
| Check current plan | [`plan-info`](references/geelark-billing-plan-info.md) |
| Upgrade plan | [`plan-upgrade`](references/geelark-billing-plan-upgrade.md) |
| Renew plan | [`plan-renew`](references/geelark-billing-plan-renew.md) |
| Buy extra minutes | [`buy-time-addon`](references/geelark-billing-buy-time-addon.md) |

## Typical Scenarios

```bash
# Check balance
geelark-cli billing balance

# View recent transactions
geelark-cli billing transaction-detail --limit 20

# See current plan
geelark-cli billing plan-info

# Upgrade plan
geelark-cli billing plan-upgrade --profiles-id "497540679501610040" --parallels-num 1 --monthly-rental-num 1

# Renew plan for 30 days
geelark-cli billing plan-renew --days 30

# Buy 2000 time add-on minutes
geelark-cli billing buy-time-addon --minutes 2000
```

## Notes

- **Balance API rate limit**: 10 requests/minute
- **Plan info API rate limit**: 1 request/minute
- **Transaction detail**: Only supports searching data within the last 3 days
- **Plan upgrade**: Only supports upgrading; downgrading must be done via the GeeLark client
- **Ensure sufficient balance** before upgrading, renewing, or buying add-ons
- **Promo codes**: Optional for `plan-upgrade`, `plan-renew`, and `buy-time-addon`

## Error Codes

| Code | Description |
|------|-------------|
| 41001 | Insufficient balance |
| 41002 | Only supports upgrade plan (downgrade via client) |
| 41003 | Promo code is invalid |

Full error codes: https://open.geelark.com/api/cloud-phone-error-codes
