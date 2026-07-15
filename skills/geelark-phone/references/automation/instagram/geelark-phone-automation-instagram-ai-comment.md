# phone automation instagram-ai-comment

Instagram AI random comment.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--random-rate <n>` | Random probability, 0-100 (required) |
| `--use-ai` | Whether to use AI for comments (default false) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Basic
geelark-cli phone automation instagram-ai-comment --id "557536075321468390" --schedule-at 1741846843 --random-rate 50

# With AI enabled
geelark-cli phone automation instagram-ai-comment --id "557536075321468390" --schedule-at 1741846843 --use-ai --random-rate 30
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
