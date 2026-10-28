# phone automation instagram-like

Like specified Instagram posts, with optional comments.

## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--links <csv>` | Comma-separated post links, max 100, each max 100 chars (required) |
| `--comments <csv>` | Comma-separated comments, max 100, each max 100 chars |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Like specified posts
geelark-cli phone automation instagram-like --id "557536075321468390" --schedule-at 1741846843 --links "https://www.instagram.com/p/xxx1/,https://www.instagram.com/p/xxx2/"

# Like and comment
geelark-cli phone automation instagram-like --id "557536075321468390" --schedule-at 1741846843 --links "https://www.instagram.com/p/xxx1/" --comments "nice post,great"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
