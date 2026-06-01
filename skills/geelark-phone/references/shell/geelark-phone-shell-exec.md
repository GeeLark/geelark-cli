# phone shell exec

Execute a shell command on a running cloud phone. Supports Android 10/12/13/14/15/16.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--cmd <text>` | Shell command to execute (required) |

## Examples

```bash
# List packages
geelark-cli phone shell exec --id "phone_id" --cmd "pm list packages"

# List files
geelark-cli phone shell exec --id "phone_id" --cmd "ls /sdcard/Download"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | bool | true=execution successful, false=execution failed |
| `output` | string | Command output |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 50001 | Cloud phone does not support shell commands |
