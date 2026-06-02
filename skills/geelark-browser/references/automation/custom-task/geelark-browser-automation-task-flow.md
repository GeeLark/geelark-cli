# browser automation task-flow

Get the list of available custom task flows for browser automation.

## Key Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |

## Examples

```bash
# List task flows
geelark-cli browser automation task-flow --page 1 --page-size 10
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `list[]` | array | Task flow list |
| `list[].id` | string | Task flow ID |
| `list[].title` | string | Task flow title |
| `list[].desc` | string | Task flow description |
| `list[].params` | array[string] | Task flow parameter field names |
