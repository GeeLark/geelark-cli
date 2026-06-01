# phone automation youtube-pub-short

YouTube publish Short.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--title <text>` | Title (max 100 chars, required) |
| `--video <text>` | Video URL (required) |
| `--same-style-url <text>` | Same style URL (max 500 chars) |
| `--same-style-voice <n>` | Same style volume, 0-100 (required; send 0 if not using same style) |
| `--original-voice <n>` | Original voice volume, 0-100 (required; send 0 if not using same style) |
| `--is-disclosure-mandatory` | Whether to force disclosure (default false) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation youtube-pub-short --id "557536075321468390" --schedule-at 1741846843 --title "My short" --video "https://material.geelark.com/a.mp4" --same-style-voice 0 --original-voice 0
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
