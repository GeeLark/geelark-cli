# phone file

File management sub-commands for cloud phones.

## Command Overview

```bash
geelark-cli phone file <command> [flags]
```

| Command | Description |
|---------|-------------|
| `upload-temp` | Upload a local file to GeeLark temp storage (3-day expiry) |
| `upload-to-phone` | Upload a file to a cloud phone's Downloads folder |
| `upload-status` | Query upload-to-phone task status |
| `keybox-upload` | Upload a keybox file for Google integrity verification |
| `keybox-result` | Query keybox upload task result |

## upload-temp

Upload a local file to GeeLark temporary storage. The CLI outputs a `resourceUrl` (valid for 3 days). For longer storage (30 days), use `library material-create` instead.

| Flag | Description |
|------|-------------|
| `--file <path>` | Local file path (required). Supported: jpg, jpeg, png, gif, bmp, webp, heif, heic, mp4, webm, xml, apk, xapk |

```bash
geelark-cli phone file upload-temp --file ./video.mp4
```

### Output

```
File uploaded successfully.
resourceUrl: https://singapore-upgrade.geelark.com/open-upload/.../xxx.jpg
Note: temporary files expire in 3 days. Use 'library material-create' for longer storage (30 days).
```

The `resourceUrl` value can be used in other commands (e.g. `upload-to-phone --url`, `keybox-upload --url`).

## upload-to-phone

Upload a file to a cloud phone's Downloads folder. Phone must be running. Supports `--url` (direct URL) or `--file` (local file, auto-uploaded to temp first).

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--url <text>` | File URL (mutually exclusive with --file) |
| `--file <path>` | Local file path (mutually exclusive with --url) |

```bash
# Via URL
geelark-cli phone file upload-to-phone --id "phone_id" --url "https://material.geelark.cn/xxx.mp4"

# Via local file
geelark-cli phone file upload-to-phone --id "phone_id" --file ./local-video.mp4
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Task ID (use with `upload-status` to query result) |

### Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |

## upload-status

Query the upload status within 1 hour of initiating the upload.

| Flag | Description |
|------|-------------|
| `--task-id <text>` | Upload task ID (required) |

```bash
geelark-cli phone file upload-status --task-id "task_id"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | 0=failed to retrieve, 1=uploading, 2=upload successful, 3=upload failed |

## keybox-upload

Upload a keybox file to pass Google's integrity verification. Only Android 12/13/15. Supports `--url` or `--file`.

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--url <text>` | Keybox file URL |
| `--file <path>` | Local keybox file path |

```bash
geelark-cli phone file keybox-upload --id "phone_id" --file ./keybox.xml
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Task ID (use with `keybox-result` to query result) |

### Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 43026 | Cloud phone system not supported |
| 60003 | Invalid file URL |

## keybox-result

Query the keybox upload task result.

| Flag | Description |
|------|-------------|
| `--task-id <text>` | Keybox upload task ID (required) |

```bash
geelark-cli phone file keybox-result --task-id "task_id"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | 0=uploading, 1=upload successful, 2=upload failed |
