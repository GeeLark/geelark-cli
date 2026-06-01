# phone net-config-set

Set cloud phone network settings including access blacklist. Max 3 blacklisted domains. Supports Android 9/10/11/12/13/15. Takes effect immediately.

## Key Flags

| Flag | Description |
|------|-------------|
| `--blacklist <csv>` | Comma-separated blacklisted domains (max 3, empty to clear) |

## Examples

```bash
# Set blacklist
geelark-cli phone net-config-set --blacklist "a.com,b.com"

# Clear blacklist
geelark-cli phone net-config-set --blacklist ""
```
