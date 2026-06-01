# phone automation instagram-edit-profile

Edit Instagram profile.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--profile-picture <csv>` | Comma-separated avatar URLs |
| `--nickname <text>` | Nickname |
| `--username <text>` | Username |
| `--biography <text>` | Biography |
| `--link-url <text>` | Link URL |
| `--link-title <text>` | Link title |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation instagram-edit-profile --id "557536075321468390" --schedule-at 1741846843 --nickname "myName" --username "myName"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
