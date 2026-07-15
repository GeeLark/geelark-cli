# phone oem customization

Customize OEM/White Label settings. Requires a plan with 50+ profiles.

## Key Flags

| Flag | Description |
|------|-------------|
| `--title <text>` | Brand title, max 64 bytes |
| `--logo <text>` | Logo URL, max 255 bytes |
| `--hide-header` | Hide the header at the top of the cloud phone |
| `--hide-sidebar` | Hide the sidebar |
| `--mirror-url <text>` | QR code/Mirror entrance URL, max 255 chars |
| `--toolbar <json>` | JSON array of toolbar settings `[{"toolBar":"...","visible":bool,"iconUrl":"..."}]` |

### Toolbar Items

`networkQuality`, `rotate`, `screenshot`, `upload`, `library`, `volumeUp`, `volumeDown`, `speedUp`, `detection`, `quality`, `restart`, `appStore`, `qcode`, `export`, `timing`, `liveStreaming`, `clear`, `teamApp`

## Examples

```bash
# Set title and logo
geelark-cli phone oem customization --title "MyBrand" --logo "https://example.com/logo.png"

# Hide header and set mirror URL
geelark-cli phone oem customization --hide-header --mirror-url "https://www.abcd.com/mirror/url"

# Hide sidebar
geelark-cli phone oem customization --hide-sidebar

# Customize toolbar
geelark-cli phone oem customization --toolbar "[{\"toolBar\":\"networkQuality\",\"visible\":false},{\"toolBar\":\"rotate\",\"visible\":false}]"

# Full customization
geelark-cli phone oem customization --title "X" --toolbar "[{\"toolBar\":\"screenshot\",\"visible\":true,\"iconUrl\":\"https://example.com/icon.svg\"}]"
```

### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

## Error Codes

| Code | Description |
|------|-------------|
| 40015 | Permission limit (need 50+ profiles) |
| 60003 | Illegal URL |
| 60004 | Invalid file format |
