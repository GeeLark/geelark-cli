---
name: geelark-group
version: 1.0.0
description: "GeeLark group management: query, create, update, and delete cloud phone groups. Use when managing cloud phone groups or organizing phones by group."
metadata:
  requires:
    bins: ["geelark-cli"]
  cliHelp: "geelark-cli group --help"
---

# group — Group Management

**CRITICAL — MUST read [`../geelark-shared/SKILL.md`](../geelark-shared/SKILL.md) first for authentication and configuration handling.**

## Command Overview

```bash
geelark-cli group <command> [flags]
```

| Command | Description | Batch |
|---------|-------------|-------|
| `list` | Query group list | — |
| `create` | Batch create groups | ✅ |
| `simple-create` | Quick create a single group | ❌ |
| `update` | Batch update groups | ✅ |
| `simple-update` | Quick update a single group | ❌ |
| `delete` | Delete groups | ✅ |

## Which Command to Use

| Want to | Command |
|---------|---------|
| View all groups | [`list`](references/geelark-group-list.md) |
| Create a group | [`simple-create`](references/geelark-group-simple-create.md) (recommended) or [`create`](references/geelark-group-create.md) (batch) |
| Modify group name/remark | [`simple-update`](references/geelark-group-simple-update.md) (recommended) or [`update`](references/geelark-group-update.md) (batch) |
| Delete a group | [`delete`](references/geelark-group-delete.md) |

## Typical Scenarios

```bash
# View all groups
geelark-cli group list --page 1 --page-size 50

# Create a group
geelark-cli group simple-create --name "marketing" --remark "Marketing team phones"

# Modify group name
geelark-cli group simple-update --id "528995439832269824" --name "marketing-v2"

# Delete a group
geelark-cli group delete --ids "528995439832269824"
```

## Notes

- **Group name must be unique**: Creating a duplicate name returns error `43031` (group name already exists)
- **"Ungrouped" is read-only**: The system default "Ungrouped" group cannot be modified or deleted; attempting so returns `43035`
- **Name**: max 50 characters; **Remark**: max 500 characters
- **Batch operations**: `create` / `update` support batch; responses include `successDetails` and `failDetails` — partial failures do not affect other items
- **Deleting a group does not delete phones**: Phones in a deleted group are moved to "Ungrouped"

## Error Codes

| Code | Description |
|------|-------------|
| 43030 | Group name is empty |
| 43031 | Group name already exists |
| 43032 | Group not found |
| 43033 | Group name not found (on query) |
| 43034 | Group remark not found (on query) |
| 43035 | Cannot operate on "Ungrouped" |

Full error codes: https://open.geelark.com/api/cloud-phone-error-codes

## Out of Scope

- Move phones into/out of groups → `geelark-cli phone move-group`
- Cloud phone management → [`geelark-phone`](../geelark-phone/SKILL.md)
