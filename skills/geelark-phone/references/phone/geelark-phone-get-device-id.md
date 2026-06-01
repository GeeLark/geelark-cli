# phone get-device-id

Get the cloud phone unique hardware device ID (Android_ID / serialno). Re-obtain after one-click new machine.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |

## Examples

```bash
geelark-cli phone get-device-id --id "phone_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `serialNum` | string | Cloud phone device ID |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
