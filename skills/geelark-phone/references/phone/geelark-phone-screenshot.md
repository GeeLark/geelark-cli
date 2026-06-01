# phone screenshot

Take a screenshot of a running cloud phone. Returns a task ID for querying the result.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |

## Examples

```bash
geelark-cli phone screenshot --id "phone_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Task ID (use with `screenshot-result`) |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
