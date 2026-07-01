# phone automation youtube-edit-profile

Edit YouTube profile.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--profile-name <text>` | Profile name (max 50 chars) |
| `--handle <text>` | Handle / identifier name (max 100 chars) |
| `--description <text>` | Description (max 1000 chars) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation youtube-edit-profile --id "557536075321468390" --schedule-at 1741846843 --profile-name "myName"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
