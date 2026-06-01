# phone screenshot-result

Query the result of a screenshot task. Valid for 30 minutes after requesting.

## Key Flags

| Flag | Description |
|------|-------------|
| `--task-id <text>` | Screenshot task ID (required) |

## Examples

```bash
geelark-cli phone screenshot-result --task-id "task_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | 0=acquisition failed, 1=in progress, 2=succeeded, 3=execution failed |
| `downloadLink` | string | Screenshot download URL (available when status=2) |
