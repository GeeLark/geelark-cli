# browser get-bookmark

Query the bookmarks of a browser environment.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Browser ID (required) |

## Examples

```bash
# Get browser bookmarks
geelark-cli browser get-bookmark --id "browser_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `browserBookmark` | object | Bookmark info |
| `browserBookmark.type` | integer | Type: 0=Not set, 1=Uploaded file, 2=Manually created |
| `browserBookmark.fileAddr` | string | Bookmark HTML file address |
| `browserBookmark.text` | string | Manually created content (formats: `Folder::Name::URL`, `Name::URL`, or `URL`, separated by `\n`) |
