# phone set-net-type

Set cloud phone network connection mode. Only supported on Android 12/13/15.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--net-type <n>` | Network type: 0=Wi-Fi, 1=Mobile (required) |

## Examples

```bash
# Set to Wi-Fi
geelark-cli phone set-net-type --id "phone_id" --net-type 0

# Set to Mobile
geelark-cli phone set-net-type --id "phone_id" --net-type 1
```

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
