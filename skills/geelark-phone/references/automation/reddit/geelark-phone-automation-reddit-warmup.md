# phone automation reddit-warmup

Reddit AI account warmup.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--keywords <csv>` | Comma-separated search keywords, max 100 (preferred) |
| `--keyword <text>` | Search keyword (deprecated, use --keywords instead) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Basic warmup
geelark-cli phone automation reddit-warmup --id "557536075321468390" --schedule-at 1741846843

# With keywords
geelark-cli phone automation reddit-warmup --id "557536075321468390" --schedule-at 1741846843 --keywords "cat,dog"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
