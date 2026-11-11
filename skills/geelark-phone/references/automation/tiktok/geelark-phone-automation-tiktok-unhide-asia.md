# phone automation tiktok-unhide-asia

Restore hidden TikTok videos (Asia region).

## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--specify-count` | Whether to restore a specific number of videos; unset = restore all |
| `--number <n>` | Number of videos to restore, range 1-999; required when `--specify-count` is set |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation tiktok-unhide-asia --id "557536075321468390" --schedule-at 1741846843
geelark-cli phone automation tiktok-unhide-asia --id "557536075321468390" --schedule-at 1741846843 --specify-count --number 10
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
