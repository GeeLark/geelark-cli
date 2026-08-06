# phone restart

Restart a cloud phone. Ensure the cloud phone startup callback has been received before calling this API.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |

## Examples

```bash
geelark-cli phone restart --id "631490227545875981"
```

## Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

## Notes

- Ensure the cloud phone startup callback has been received before calling restart
- This API restarts a single phone (use `--id`, not `--ids`)

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 43005 | Cloud phone is executing a task |
