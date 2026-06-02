# browser clear-cache

Clear local cache of browser environments. Ensure browsers are closed first.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Comma-separated browser IDs (required) |
| `--type <csv>` | Comma-separated cache types (required) |

### Cache Types

| Type | Description |
|------|-------------|
| `local_storage` | Local storage |
| `indexeddb` | IndexedDB |
| `extension_cache` | Extension cache |
| `cookie` | Cookies |
| `history` | Browsing history |
| `image_file` | Image files |

## Examples

```bash
# Clear cookie and history
geelark-cli browser clear-cache --ids "id1,id2" --type "cookie,history"

# Clear all cache types
geelark-cli browser clear-cache --ids "id1" --type "local_storage,indexeddb,extension_cache,cookie,history,image_file"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `success_count` | integer | Number of successfully cleared browsers |
| `error_count` | integer | Number of failed browsers |
| `error_info` | array | Error details |
