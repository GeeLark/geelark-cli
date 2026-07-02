# phone new-one

One-click new machine — reset cloud phone identity. Applications and data will be cleared.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--no-change-brand` | Do NOT randomize brand/model (keep unchanged) |
| `--keep-net-type` | Preserve network connection type |
| `--keep-phone-number` | Preserve phone number |
| `--keep-region` | Preserve region (otherwise follows proxy) |
| `--keep-language` | Preserve language (otherwise defaults to English) |
| `--mobile-type <text>` | Change mobile type: Android 9, 10, 11, 12, 13, 14, 15, 16 |

## Examples

```bash
# Full reset
geelark-cli phone new-one --id "phone_id"

# Reset but keep some settings
geelark-cli phone new-one --id "phone_id" --keep-region --keep-language --keep-phone-number

# Reset but keep brand
geelark-cli phone new-one --id "phone_id" --no-change-brand

# Reset and change Android version
geelark-cli phone new-one --id "phone_id" --mobile-type "Android 16"
```

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 43005 | Cloud phone is executing a task |
| 43006 | Cloud phone is occupied by remote |
| 43015 | Cloud phone does not support new-one |
| 43038 | Device brand/model deleted |
| 44004 | Daily creation limit reached |
| 45004 | Proxy detect failed |
