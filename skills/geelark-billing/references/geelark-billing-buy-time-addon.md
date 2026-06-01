# billing buy-time-addon

Buy time add-on minutes. Ensure sufficient balance before purchasing.

## Key Flags

| Flag | Description |
|------|-------------|
| `--minutes <n>` | Minutes to buy (required). Available: 2000, 5000, 10000, 20000, 50000, 100000, 200000, 500000, 1000000, 2000000, 5000000, 10000000 |
| `--promo-code <text>` | Promo code (optional) |

## Examples

```bash
# Buy 2000 minutes
geelark-cli billing buy-time-addon --minutes 2000

# Buy 10000 minutes with promo code
geelark-cli billing buy-time-addon --minutes 10000 --promo-code "GL666"
```

## Error Codes

| Code | Description |
|------|-------------|
| 41001 | Insufficient balance |
| 41003 | Promo code is invalid |
