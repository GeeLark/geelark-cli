# browser automation task-cancel

Cancel a browser automation task. Can cancel while running or waiting to be executed.

## Key Flags

| Flag | Description |
|------|-------------|
| `--task-id <text>` | Task ID (required) |

## Examples

```bash
# Cancel a task
geelark-cli browser automation task-cancel --task-id "497652752864775437"
```

## Response Fields

> The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

On success, `code` is `0` and `msg` is `"success"`. No additional `data` fields.

## Error Codes

| Code | Description |
|------|-------------|
| 48001 | Task status does not allow operation |
| 48005 | Only the creator is allowed to operate |
