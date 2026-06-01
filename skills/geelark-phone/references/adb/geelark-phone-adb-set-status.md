# phone adb set-status

Enable or disable ADB on cloud phones. Supports Android 9/11/12/13/14/15/16. Phone must be started first. Wait ~3 seconds after enabling before retrieving connection info.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs (required) |
| `--open` | Enable ADB (default true). Use `--open=false` to disable |

## Examples

```bash
# Enable ADB
geelark-cli phone adb set-status --ids "id1,id2" --open

# Disable ADB
geelark-cli phone adb set-status --ids "id1" --open=false
```
