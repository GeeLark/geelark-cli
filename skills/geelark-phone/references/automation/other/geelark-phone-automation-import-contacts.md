# phone automation import-contacts

Batch import contacts to cloud phone.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--contacts <json>` | Contacts array as JSON string (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

### Contacts JSON Format

```json
[
  {"firstName": "jay", "mobile": "13288888888"},
  {"firstName": "tom", "mobile": "13299999999"}
]
```

## Example

```bash
geelark-cli phone automation import-contacts --id "557536075321468390" --schedule-at 1741846843 --contacts '[{"firstName":"jay","mobile":"13288888888"}]'
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
