---
name: geelark-proxy
version: 1.0.0
description: "GeeLark proxy management: list, add, update, delete, and check proxies. Use when managing proxy configurations for cloud phones."
metadata:
  requires:
    bins: ["geelark-cli"]
  cliHelp: "geelark-cli proxy --help"
---

# proxy — Proxy Management

**CRITICAL — MUST read [`../geelark-shared/SKILL.md`](../geelark-shared/SKILL.md) first for authentication and configuration handling.**

## Command Overview

```bash
geelark-cli proxy <command> [flags]
```

| Command | Description | Batch |
|---------|-------------|-------|
| `list` | List all proxies | — |
| `add` | Batch add proxies | ✅ |
| `simple-add` | Quick add a single proxy | ❌ |
| `update` | Batch update proxies | ✅ |
| `simple-update` | Quick update a single proxy | ❌ |
| `delete` | Delete proxies | ✅ |
| `check` | Check/detect a proxy | ❌ |

## Which Command to Use

| Want to | Command |
|---------|---------|
| View all proxies | [`list`](references/geelark-proxy-list.md) |
| Add a proxy | [`simple-add`](references/geelark-proxy-simple-add.md) (recommended) or [`add`](references/geelark-proxy-add.md) (batch) |
| Modify a proxy | [`simple-update`](references/geelark-proxy-simple-update.md) (recommended) or [`update`](references/geelark-proxy-update.md) (batch) |
| Delete proxies | [`delete`](references/geelark-proxy-delete.md) |
| Test proxy connectivity | [`check`](references/geelark-proxy-check.md) |

## Typical Scenarios

```bash
# List proxies
geelark-cli proxy list --page 1 --page-size 10

# Add a proxy
geelark-cli proxy simple-add --scheme socks5 --server 192.3.8.1 --port 8000

# Update a proxy
geelark-cli proxy simple-update --id "proxy_id" --server 10.0.0.1 --port 9000

# Delete proxies
geelark-cli proxy delete --ids "id1,id2"

# Check proxy connectivity
geelark-cli proxy check --type socks5 --server 185.162.130.86 --port 10000
```

## Notes

- **Duplicate proxies are ignored**: Adding a proxy that already exists will not create a duplicate
- **Proxy types**: `socks5`, `http`, `https`
- **Detection channels**: `1` = IP-API, `2` = IP2Location (default)
- **Batch limit**: Up to 100 proxy items per batch request
- **Proxy in use cannot be deleted**: A proxy bound to a cloud phone returns error `40010`

## Error Codes

| Code | Description |
|------|-------------|
| 40005 | Proxy not found |
| 40010 | Proxy is bound to an environment (cannot delete) |
| 45003 | Proxy not allowed |
| 45004 | Check proxy failed |
| 45007 | Proxy already exists |
| 45008 | Proxy type not allowed |

Full error codes: https://open.geelark.com/api/cloud-phone-error-codes
