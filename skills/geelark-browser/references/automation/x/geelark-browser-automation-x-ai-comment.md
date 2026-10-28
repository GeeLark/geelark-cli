# browser automation x-ai-comment

Create an X (Twitter) automation task to post AI random comments.

## Key Flags

| Flag | Description |
|------|-------------|
| `--eid <text>` | Environment ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--use-ai <n>` | Use AI: 1=AI (Pro only), 2=manual comment (required) |
| `--comment <text>` | Comment content, max 300 chars (required when --use-ai=2) |
| `--links <csv>` | Comma-separated specified links, max 100, each max 100 chars |
| `--comment-probability <n>` | Comment probability, 0-100, default 30 |
| `--search-keywords <csv>` | Comma-separated search keywords, max 100, each max 100 chars |
| `--like-video` | Whether to like, defaults to false |
| `--image-path <text>` | Comment image local file path, max 500 chars (effective when --use-ai=2) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Manual comment
geelark-cli browser automation x-ai-comment --eid "557536075321468390" --schedule-at 1741846843 --use-ai 2 --comment "test"

# Manual comment with probability and like
geelark-cli browser automation x-ai-comment --eid "557536075321468390" --schedule-at 1741846843 --use-ai 2 --comment "test" --comment-probability 50 --like-video

# AI comment (Pro only)
geelark-cli browser automation x-ai-comment --eid "557536075321468390" --schedule-at 1741846843 --use-ai 1
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
