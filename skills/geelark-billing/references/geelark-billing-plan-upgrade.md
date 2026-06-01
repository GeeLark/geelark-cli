# billing plan-upgrade

Upgrade subscription plan. Only supports upgrading; downgrading must be done via the GeeLark client. Use `plan-list` to get available `profilesId` values.

## Key Flags

| Flag | Description |
|------|-------------|
| `--profiles-id <text>` | Plan profiles ID from plan-list (required) |
| `--parallels-num <n>` | Parallels number, must be >= current plan (default 1) |
| `--monthly-rental-num <n>` | Monthly rental number, must be >= current plan (default 0) |
| `--days <n>` | Renewal duration: 30/90/180/360 days (required when plan is expired) |
| `--promo-code <text>` | Promo code (optional) |

## Examples

```bash
# Upgrade plan
geelark-cli billing plan-upgrade --profiles-id "497540679501610040" --parallels-num 1 --monthly-rental-num 1

# Upgrade with renewal duration
geelark-cli billing plan-upgrade --profiles-id "512719311391949750" --parallels-num 3 --monthly-rental-num 5 --days 30

# Upgrade with promo code
geelark-cli billing plan-upgrade --profiles-id "id" --parallels-num 1 --monthly-rental-num 1 --promo-code "GL666"
```

## Notes

- **Ensure sufficient balance** before upgrading
- **Only supports upgrading**; downgrading must be done via the GeeLark client
- Use `plan-list` to get valid `profilesId` values

## Error Codes

| Code | Description |
|------|-------------|
| 41001 | Insufficient balance |
| 41002 | Only supports upgrade plan (downgrade via client) |
| 41003 | Promo code is invalid |
