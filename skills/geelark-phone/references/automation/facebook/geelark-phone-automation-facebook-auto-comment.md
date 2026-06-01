# phone automation facebook-auto-comment

Facebook auto comment.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--post-address <text>` | Post address (max 128 chars, required) |
| `--comment <csv>` | Comma-separated comments, max 10 (required) |
| `--keyword <csv>` | Comma-separated keywords, max 10 (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation facebook-auto-comment --id "557536075321468390" --schedule-at 1741846843 --post-address "https://abc.com" --comment "test1,test2" --keyword "k1,k2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
