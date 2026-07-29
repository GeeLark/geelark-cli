# phone set-net-type

Set cloud phone network connection mode. Only supported on Android 12-16.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--net-type <n>` | Network type: 0=Wi-Fi, 1=Mobile (required) |
| `--wifi-id <text>` | Wi-Fi name, max 16 chars |

## Examples

```bash
# Set to Wi-Fi
geelark-cli phone set-net-type --id "phone_id" --net-type 0

# Set to Mobile
geelark-cli phone set-net-type --id "phone_id" --net-type 1

# Set to Wi-Fi with custom name
geelark-cli phone set-net-type --id "phone_id" --net-type 0 --wifi-id "TP-link"
```

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
