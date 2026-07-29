# proxy group-add

Create a proxy group. The group name must be unique within the team (max 50 chars).

## Key Flags

| Flag | Description |
|------|-------------|
| `--name <text>` | Proxy group name, max 50 chars, must be unique (required) |

## Examples

```bash
geelark-cli proxy group-add --name "Business Group A"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Created proxy group ID |

## Error Codes

| Code | Description |
|------|-------------|
| 40004 | Invalid argument |
| 45010 | Proxy group already exists |
| 45011 | Proxy group name is not allowed |
