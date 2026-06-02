# browser automation task-detail

Get detailed information of a browser automation task by task ID, including task logs.

## Key Flags

| Flag | Description |
|------|-------------|
| `--task-id <text>` | Task ID (required) |

## Examples

```bash
# Get task details
geelark-cli browser automation task-detail --task-id "497652752864775437"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Task ID |
| `eid` | string | Environment ID |
| `name` | string | Task name |
| `remark` | string | Task remark |
| `serialName` | string | Environment name |
| `status` | integer | Status: 1=Waiting, 2=Executing, 3=Completed, 4=Failed, 7=Cancelled |
| `startAt` | integer | Start time (seconds timestamp) |
| `finishAt` | integer | End time (seconds timestamp) |
| `cost` | integer | Duration in seconds |
| `resultCode` | integer | Result code (0=success) |
| `resultDesc` | string | Result description |
| `scheduleAt` | integer | Scheduled execution time (seconds timestamp) |
| `logs` | array[string] | Task execution logs |
