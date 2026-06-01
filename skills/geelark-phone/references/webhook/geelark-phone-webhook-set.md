# phone webhook set

Set the webhook callback URL for receiving notifications.

## Key Flags

| Flag | Description |
|------|-------------|
| `--url <text>` | Callback URL (required) |

## Examples

```bash
geelark-cli phone webhook set --url "https://example.com/callback"
```

### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).
