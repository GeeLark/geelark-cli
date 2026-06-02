# browser set-bookmark

Set/update bookmarks for a browser environment. Bookmarks are applied automatically when the browser starts.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON bookmark data (required) |

### JSON Parameters (in --data)

| Parameter | Required | Type | Description |
|-----------|----------|------|-------------|
| `id` | Yes | string | Browser ID |
| `browserBookmark` | Yes | object | Bookmark configuration |
| `browserBookmark.type` | Yes | integer | Type: 0=Not set, 1=Uploaded file, 2=Manually created |
| `browserBookmark.fileAddr` | No | string | Bookmark HTML file address (for type=1) |
| `browserBookmark.text` | No | string | Manually created content (for type=2). Formats: `Folder::Name::URL`, `Name::URL`, or `URL`, separated by `\n` |

## Examples

```bash
# Set manual bookmarks
geelark-cli browser set-bookmark --data '{"id":"browser_id","browserBookmark":{"type":2,"fileAddr":"","text":"http://a.com\nhttp://b.com"}}'

# Set bookmarks with folders
geelark-cli browser set-bookmark --data '{"id":"browser_id","browserBookmark":{"type":2,"text":"Social::Google::https://google.com\nhttps://github.com"}}'
```

## Response Fields

> The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

On success, `code` is `0` and `msg` is `"success"`. No additional `data` fields.
