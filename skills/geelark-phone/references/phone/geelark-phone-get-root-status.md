# phone get-root-status

Query whether root is enabled on cloud phones. Batch results are returned per phone.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs, 1-100 (required) |

## Examples

```bash
geelark-cli phone get-root-status --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `items[]` | array | Per-phone result |
| `items[].id` | string | Cloud phone ID |
| `items[].code` | integer | Result code (0=success) |
| `items[].msg` | string | Result message |
| `items[].open` | bool | Whether root is enabled |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone not found |

Related: enable/disable root via [`set-root`](geelark-phone-set-root.md).
