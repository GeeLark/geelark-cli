# phone patch uninstall

Uninstall an Xposed patch from a cloud phone. Supports Android 12/13/14/15/16.

- Android 12/13/15: pass `--patch-id`.
- Android 14/16: pass `--patch-name` (must match the one used at install time).

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--patch-id <text>` | Patch ID, for Android 12/13/15 |
| `--patch-name <text>` | Patch name used at install time, for Android 14/16 |

## Examples

```bash
# Android 12/13/15
geelark-cli phone patch uninstall --id "532930887177274368" --patch-id "sdktest"

# Android 14/16
geelark-cli phone patch uninstall --id "643654639442503208" --patch-name "hookplugin-debug"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | Uninstall status: 0=failed, 1=success |
