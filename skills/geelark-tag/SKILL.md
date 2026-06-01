---
name: geelark-tag
version: 1.0.0
description: "GeeLark tag management: list, create, update, and delete cloud phone tags. Use when managing tags for organizing cloud phones."
metadata:
  requires:
    bins: ["geelark-cli"]
  cliHelp: "geelark-cli tag --help"
---

# tag — Tag Management

**CRITICAL — MUST read [`../geelark-shared/SKILL.md`](../geelark-shared/SKILL.md) first for authentication and configuration handling.**

## Command Overview

```bash
geelark-cli tag <command> [flags]
```

| Command | Description | Batch |
|---------|-------------|-------|
| `list` | List tags | — |
| `create` | Batch create tags | ✅ |
| `simple-create` | Quick create a single tag | ❌ |
| `update` | Batch update tags | ✅ |
| `simple-update` | Quick update a single tag | ❌ |
| `delete` | Delete tags | ✅ |

## Which Command to Use

| Want to | Command |
|---------|---------|
| View all tags | [`list`](references/geelark-tag-list.md) |
| Create a tag | [`simple-create`](references/geelark-tag-simple-create.md) (recommended) or [`create`](references/geelark-tag-create.md) (batch) |
| Modify tag name/color | [`simple-update`](references/geelark-tag-simple-update.md) (recommended) or [`update`](references/geelark-tag-update.md) (batch) |
| Delete tags | [`delete`](references/geelark-tag-delete.md) |

## Typical Scenarios

```bash
# List tags
geelark-cli tag list --page 1 --page-size 50

# Create a tag
geelark-cli tag simple-create --name "marketing" --color red

# Update a tag
geelark-cli tag simple-update --id "tag_id" --name "sales" --color blue

# Delete tags
geelark-cli tag delete --ids "id1,id2"
```

## Notes

- **Tag name must be unique**: Creating a duplicate name returns error `43021`
- **Tag name**: max 30 characters
- **Tag colors**: `white` (default), `red`, `blue`, `green`, `yellow`, `purple`
- **Batch operations**: `create` / `update` support batch; responses include `successDetails` and `failDetails`

## Error Codes

| Code | Description |
|------|-------------|
| 43020 | Tag name is empty |
| 43021 | Tag name already exists |
| 43022 | Tag does not exist |
| 43023 | Tag color not supported |
| 43024 | Tag name not found (on query) |

Full error codes: https://open.geelark.com/api/cloud-phone-error-codes
