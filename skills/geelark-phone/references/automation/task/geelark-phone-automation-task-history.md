# phone automation task-history

Batch query task history. Returns all tasks scheduled within the past 7 days.


## Flags

| Flag | Description |
|------|-------------|
| `--size <n>` | Number of records per page, max 100 |
| `--last-id <text>` | Last item ID from previous page for pagination |
| `--ids <csv>` | Comma-separated task IDs, max 100 |

## Example

```bash
geelark-cli phone automation task-history --size 10
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

### Error Codes

Task failure codes are extensive (20002-21001, 29992-29999). See full list at https://open.geelark.com/api/cloud-phone-error-codes
