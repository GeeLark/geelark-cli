# phone automation add-custom-task

Create a custom automation task using a task flow ID. First call `task-flow-list` to get available task flows.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--flow-id <text>` | Task flow ID (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |
| `--param-map <json>` | Task flow parameters as JSON string. Types: String=string, Batch text=array[string], Number=number, Boolean=bool, File=array[string] |

## Example

```bash
geelark-cli phone automation add-custom-task --id "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885"
geelark-cli phone automation add-custom-task --id "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885" --param-map '{"Title":"video","Video":["https://material.geelark.com/a.mp4"]}'
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
