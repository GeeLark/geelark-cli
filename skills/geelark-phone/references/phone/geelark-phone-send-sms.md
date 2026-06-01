# phone send-sms

Send SMS to a cloud phone. Phone must be started first. Supports Android 12/13/14/15.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--phone-number <text>` | Phone number with country code (required) |
| `--text <text>` | SMS content (required) |

## Examples

```bash
geelark-cli phone send-sms --id "phone_id" --phone-number "+17723504471" --text "your code: 6666"
```

## Error Codes

| Code | Description |
|------|-------------|
| 52001 | This type of cloud phone does not support sending SMS |
