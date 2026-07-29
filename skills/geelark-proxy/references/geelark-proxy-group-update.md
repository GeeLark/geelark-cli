# proxy group-update

Update a proxy group name. The new name must be unique within the team (max 50 chars).

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Proxy group ID (required) |
| `--name <text>` | New proxy group name, max 50 chars (required) |

## Examples

```bash
geelark-cli proxy group-update --id "123456789012345678" --name "New Group Name"
```

## Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

## Error Codes

| Code | Description |
|------|-------------|
| 40004 | Invalid argument |
| 45009 | Proxy group not found |
| 45010 | Proxy group already exists |
| 45011 | Proxy group name is not allowed |
