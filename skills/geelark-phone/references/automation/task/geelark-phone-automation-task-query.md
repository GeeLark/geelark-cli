# phone automation task-query

Query cloud phone tasks by IDs.


## Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Comma-separated task IDs, max 100 (required) |

## Example

```bash
geelark-cli phone automation task-query --ids "id1,id2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total number of tasks |
| `items[]` | array[Task] | Task array |

### Task Object

| Field | Type | Description |
|-------|------|-------------|
| `items[].id` | string | Task ID |
| `items[].planName` | string | Task plan name |
| `items[].taskType` | integer | Task type: 1=TikTok video, 2=TikTok warmup, 3=TikTok image set, 4=TikTok login, 6=TikTok profile edit, 42=Custom |
| `items[].serialName` | string | Cloud phone name |
| `items[].envId` | string | Cloud phone ID |
| `items[].scheduleAt` | integer | Scheduled time (seconds timestamp) |
| `items[].status` | integer | 1=Waiting, 2=In progress, 3=Completed, 4=Failed, 7=Cancelled |
| `items[].failCode` | integer | Failure code (present when status=4) |
| `items[].failDesc` | string | Failure reason (present when status=4) |
| `items[].cost` | integer | Duration in seconds (completed/failed) |
| `items[].shareLink` | string | Share link |

### Error Codes

Task failure codes are extensive (20002-21001, 29992-29999). See [Task Failure Codes](#task-failure-codes-reference) below or full list at https://open.geelark.com/api/cloud-phone-error-codes

#### Task Failure Codes Reference

| Code | Description |
|------|-------------|
| 20002 | Machine is performing other tasks |
| 20003 | Execution timeout |
| 20005 | Task canceled |
| 20007 | Unsupported task type |
| 20100 | No network connection |
| 20129 | Device offline |
| 20130 | Account password is wrong |
| 20136 | Account blocked |
| 20200-20266 | Video/media processing errors |
| 20300-20340 | Registration errors |
| 29992 | Environment startup failed |
| 29993 | System error |
| 29994 | Network unstable |
| 29995 | Currently unavailable; maintenance in progress |
| 29996 | Proxy detection failed |
| 29997 | Insufficient balance |
| 29998 | Cloud phone has been deleted |
| 29999 | Unknown error |
