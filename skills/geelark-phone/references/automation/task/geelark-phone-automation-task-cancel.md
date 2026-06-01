# phone automation task-cancel

Cancel tasks that are in Waiting or In progress status.


## Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Comma-separated task IDs, max 100 (required) |

## Example

```bash
geelark-cli phone automation task-cancel --ids "id1,id2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total number processed |
| `successAmount` | integer | Number of successfully processed tasks |
| `failAmount` | integer | Number of failed tasks |
| `failDetails[]` | array | Details of failed tasks (only present if some fail) |
| `failDetails[].id` | string | Task ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

### Error Codes

| Code | Description |
|------|-------------|
| 48001 | Task status does not allow cancellation |
| 40000 | Unknown error |
