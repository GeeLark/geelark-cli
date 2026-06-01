# phone automation task-restart

Retry failed or cancelled tasks. A task can be retried up to 5 times. Tasks created via API do not auto-retry.


## Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Comma-separated task IDs (required) |

## Example

```bash
geelark-cli phone automation task-restart --ids "id1,id2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total number of tasks processed |
| `successAmount` | integer | Number of tasks processed successfully |
| `failAmount` | integer | Number of tasks failed to process |
| `failDetails[]` | array | Details of failed tasks (only present if some fail) |
| `failDetails[].id` | string | Task ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

### Error Codes

| Code | Description |
|------|-------------|
| 40005 | Environment has been deleted |
| 48000 | Task retry limit reached |
| 48001 | Task status does not allow retry |
| 48002 | Task does not exist |
| 48003 | The task resource has expired |
