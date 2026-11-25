# billing gproxy-buy

Buy GProxy traffic, charged from the team wallet balance. Ensure sufficient balance before purchasing.

## Key Flags

| Flag | Description |
|------|-------------|
| `--traffic <n>` | Traffic to buy in GB (required). Must be a valid tier: 5/10/20/50/100/150/200/300/400/500/600/700/800/900/1000/2000/3000/4000/5000 |
| `--promo-code <text>` | Promo code (optional) |

## Examples

```bash
# Buy 10 GB of GProxy traffic
geelark-cli billing gproxy-buy --traffic 10

# Buy 10 GB with promo code
geelark-cli billing gproxy-buy --traffic 10 --promo-code "GeeLark666"
```

## Error Codes

| Code | Description |
|------|-------------|
| 41001 | Insufficient balance |
| 41003 | Promo code is invalid |
