# phone automation task-detail

Query cloud phone task detail with logs.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Task ID (required) |
| `--search-after <json>` | Log pagination parameter as JSON array |

## Example

```bash
geelark-cli phone automation task-detail --id "1234567898"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Task ID |
| `planName` | string | Task plan name |
| `taskType` | integer | Task type: 1=TikTok video, 2=TikTok warmup, 3=TikTok image set, 4=TikTok login, 6=TikTok profile edit, 42=Custom |
| `serialName` | string | Cloud phone name |
| `envId` | string | Cloud phone ID |
| `scheduleAt` | integer | Scheduled time (seconds timestamp) |
| `status` | integer | 1=Waiting, 2=In progress, 3=Completed, 4=Failed, 7=Cancelled |
| `failCode` | integer | Failure code (present when status=4) |
| `failDesc` | string | Failure reason (present when status=4) |
| `cost` | integer | Duration in seconds (completed/failed) |
| `resultImages[]` | array[string] | Screenshots taken when task completes or fails |
| `logs[]` | array[string] | Task logs (timezone: UTC) |
| `searchAfter[]` | array[integer] | Log pagination parameter for next page |
| `logContinue` | bool | Whether more logs are available |

### Pagination

If `logContinue` is `true`, pass the `searchAfter` value from the response as `--search-after` in the next request to fetch more logs.
