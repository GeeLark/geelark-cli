# phone automation task-flow-list

Query custom task flows.


## Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (required) |
| `--page-size <n>` | Page size, max 100 (required) |

## Example

```bash
geelark-cli phone automation task-flow-list --page 1 --page-size 10
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total number of items |
| `page` | integer | Page number |
| `pageSize` | integer | Page size |
| `items[]` | array[TaskFlow] | Task flow array |
| `items[].id` | string | Task flow ID |
| `items[].title` | string | Task flow title |
| `items[].desc` | string | Task flow description |
| `items[].params[]` | array[string] | Parameter field names |
