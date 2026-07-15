# phone automation tiktok-edit-profile

TikTok profile editing.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--avatar <text>` | Avatar URL (1:1 aspect ratio recommended) |
| `--nick-name <text>` | Nickname (max 30 chars) |
| `--bio <text>` | Bio (max 160 chars) |
| `--site <text>` | Website URL (must start with http/https) |
| `--username <text>` | Username |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation tiktok-edit-profile --id "557536075321468390" --schedule-at 1741846843 --nick-name "myName" --bio "Hello"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
