# group simple-create

Quick create a single group using flat flags instead of JSON. For batch creation, use `create`.

## Key Flags

| Flag | Description |
|------|-------------|
| `--name <text>` | Group name, max 50 characters (required) |
| `--remark <text>` | Group remark, max 500 characters (optional) |

## Examples

```bash
# Name only
geelark-cli group simple-create --name "marketing"

# Name and remark
geelark-cli group simple-create --name "marketing" --remark "Marketing team phones"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../geelark-shared/SKILL.md#api-response-format). Response structure is identical to [`create`](geelark-group-create.md).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total request count (always 1 for simple-create) |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `successDetails[]` | array | Successful details |
| `successDetails[].id` | string | Newly created group ID |
| `successDetails[].name` | string | Group name |
| `failDetails[]` | array | Failed details |
| `failDetails[].code` | integer | Error code |
| `failDetails[].name` | string | Failed group name |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 43030 | Group name is empty |
| 43031 | Group name already exists |
