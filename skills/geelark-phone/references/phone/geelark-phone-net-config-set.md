# phone net-config-set

Set cloud phone network settings including access blacklist and proxy direct whitelist.

- Access blacklist: max 3 domains, blocks cloud phone access to these domains, supports Android 9/10/11/12/13/15. Takes effect immediately.
- Proxy direct whitelist: max 3 domains that bypass the proxy and connect directly (saves traffic), supports Android 12/13/14/15/16. Subdomains are automatically included (e.g. `example.com` covers subdomains, no wildcard needed).

Each list is only updated when its flag is explicitly set; pass an empty string to clear it.

## Key Flags

| Flag | Description |
|------|-------------|
| `--blacklist <csv>` | Comma-separated blacklisted domains (max 3, empty to clear; omit to leave unchanged) |
| `--proxy-direct <csv>` | Comma-separated proxy direct whitelist domains (max 3, subdomains included, empty to clear; omit to leave unchanged) |

## Examples

```bash
# Set blacklist
geelark-cli phone net-config-set --blacklist "a.com,b.com"

# Clear blacklist
geelark-cli phone net-config-set --blacklist ""

# Set proxy direct whitelist
geelark-cli phone net-config-set --proxy-direct "example.com"

# Update both
geelark-cli phone net-config-set --blacklist "a.com" --proxy-direct "b.com"
```
