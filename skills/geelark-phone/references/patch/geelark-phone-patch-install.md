# phone patch install

Install an Xposed patch on a cloud phone. Supports Android 12/13/14/15/16.

- Android 12/13/15: original install mode, `--apk-url` supports APK and XAPK; `--patch-name` / `--target-app` are not needed.
- Android 14/16: apmt install mode, `--apk-url` supports APK only; `--patch-name` and `--target-app` are required.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--apk-url <text>` | Patch install file URL (required) |
| `--patch-name <text>` | Patch name, 1-64 chars, letters/digits/underscore/hyphen only, no leading hyphen (required for Android 14/16) |
| `--target-app <text>` | Hook target app package name, max 128 chars; `system` means system hook (required for Android 14/16) |

## Examples

```bash
# Android 12/13/15
geelark-cli phone patch install --id "493188646250219877" --apk-url "https://material.geelark.com/xxx.apk"

# Android 14/16
geelark-cli phone patch install --id "493188646250219877" --apk-url "https://material.geelark.com/xxx.apk" --patch-name "geelarkpatch" --target-app "com.abc"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Task ID, used to query the upload and install status via `patch status` |
