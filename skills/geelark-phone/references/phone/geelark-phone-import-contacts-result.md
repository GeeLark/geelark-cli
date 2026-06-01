# phone import-contacts-result

Query the result of a contact import task. Valid for 1 hour after creation.

## Key Flags

| Flag | Description |
|------|-------------|
| `--task-id <text>` | Import contacts task ID (required) |

## Examples

```bash
geelark-cli phone import-contacts-result --task-id "task_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | Task status: 1=in progress, 2=successful, 3=failed |
