# browser automation task-search

Query the list of browser automation tasks with optional filters.

## Key Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |
| `--task-ids <csv>` | Comma-separated task IDs |

## Examples

```bash
# List tasks
geelark-cli browser automation task-search --page 1 --page-size 10

# Filter by task IDs
geelark-cli browser automation task-search --task-ids "497652752864775437"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `list[]` | array | Task list |
| `list[].id` | string | Task ID |
| `list[].eid` | string | Environment ID |
| `list[].name` | string | Task name |
| `list[].remark` | string | Task remark |
| `list[].serialName` | string | Environment name |
| `list[].status` | integer | Status: 1=Waiting, 2=Executing, 3=Completed, 4=Failed, 7=Cancelled |
| `list[].startAt` | integer | Start time (seconds timestamp) |
| `list[].finishAt` | integer | End time (seconds timestamp) |
| `list[].cost` | integer | Duration in seconds |
| `list[].resultCode` | integer | Result code (0=success) |
| `list[].resultDesc` | string | Result description |
| `list[].scheduleAt` | integer | Scheduled execution time (seconds timestamp) |
