# phone automation instagram-pub-reels-images

Instagram publish Reels image set.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--description <text>` | Caption (max 2200 chars, required) |
| `--image <csv>` | Comma-separated image URLs, max 10 (required) |
| `--same-style-url <text>` | Same style URL |
| `--ai-tag` | AI tag (default false) |
| `--publish-post` | Whether to also publish as POST (default false) |
| `--need-share-link` | Whether to retrieve sharing link (default false) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation instagram-pub-reels-images --id "557536075321468390" --schedule-at 1741846843 --description "My images" --image "https://material.geelark.com/a.jpg"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
