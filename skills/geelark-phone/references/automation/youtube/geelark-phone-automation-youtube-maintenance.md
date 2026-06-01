# phone automation youtube-maintenance

YouTube account maintenance (browse videos).


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--browse-video-num <n>` | Number of videos to browse, 1-100 (required) |
| `--keyword <csv>` | Comma-separated keywords, max 10 (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation youtube-maintenance --id "557536075321468390" --schedule-at 1741846843 --browse-video-num 10 --keyword "k1,k2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
