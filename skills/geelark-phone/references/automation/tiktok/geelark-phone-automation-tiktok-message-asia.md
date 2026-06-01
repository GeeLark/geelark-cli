# phone automation tiktok-message-asia

Send private message on TikTok — Asia region.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--usernames <csv>` | Comma-separated usernames (required) |
| `--content <text>` | Message content, max 6000 chars (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation tiktok-message-asia --id "557536075321468390" --schedule-at 1741846843 --usernames "user1" --content "Hello!"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
