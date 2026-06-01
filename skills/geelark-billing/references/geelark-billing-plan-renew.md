# billing plan-renew

Renew the current subscription plan. Ensure sufficient balance before renewing.

## Key Flags

| Flag | Description |
|------|-------------|
| `--days <n>` | Renewal duration: 30, 90, 180, or 360 days (required) |
| `--promo-code <text>` | Promo code (optional) |

## Examples

```bash
# Renew for 30 days
geelark-cli billing plan-renew --days 30

# Renew for 90 days with promo code
geelark-cli billing plan-renew --days 90 --promo-code "GL666"
```

## Error Codes

| Code | Description |
|------|-------------|
| 41001 | Insufficient balance |
| 41003 | Promo code is invalid |
