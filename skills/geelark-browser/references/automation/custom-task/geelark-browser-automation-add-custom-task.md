# browser automation add-custom-task

Create a custom browser automation task using a task flow ID. First call `task-flow` to get available task flows, then create a task with the flow ID.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--flow-id <text>` | Task flow ID (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |
| `--param-map <json>` | Task flow parameters as JSON string |

### paramMap Parameter Types

| Flow Parameter Type | JSON Type | Example |
|---------------------|-----------|---------|
| String | string | `"video"` |
| Batch text | array[string] | `["text1","text2"]` |
| Number | number | `42` |
| Boolean | bool | `true` |
| File | array[string] | `["https://material.geelark.cn/a.mp4"]` |

## Examples

```bash
# Create custom task with parameters
geelark-cli browser automation add-custom-task --eid "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885" --param-map '{"Title":"video","Desc":"this is video"}'

# Create with file parameter
geelark-cli browser automation add-custom-task --eid "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885" --param-map '{"Title":"video","Video":["https://material.geelark.cn/a.mp4"]}'
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |

## Error Codes

| Code | Description |
|------|-------------|
| 43028 | User does not have permission for this environment group |
| 43027 | Environment not supported |
| 46002 | Package expired, member unavailable |
| 46003 | Package expired, environment unavailable |
