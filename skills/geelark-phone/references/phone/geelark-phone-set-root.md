# phone set-root

Enable or disable root on cloud phones. Phone must be started first. Supports Android 12/13/14/15/16.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs (required) |
| `--open` | Enable root (default true). Use `--open=false` to disable |

## Examples

```bash
# Enable root
geelark-cli phone set-root --ids "id1,id2" --open

# Disable root
geelark-cli phone set-root --ids "id1" --open=false
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `items[]` | array | Per-phone result |
| `items[].id` | string | Cloud phone ID |
| `items[].code` | integer | Result code (0=success) |
| `items[].msg` | string | Result message |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 43016 | Cloud phone does not support root |
