# phone ui-dump

Dump the UI layout of a cloud phone for automation element locating. Async task — returns a `taskId`; query the result with [`ui-dump-result`](geelark-phone-ui-dump-result.md) within 30 minutes.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required). Phone must be running |

## Examples

```bash
geelark-cli phone ui-dump --id "phone_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | UI Dump task ID, used to query the result |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |

A UI Dump completion callback (webhook event type 16) is also available; see the [webhook set](../webhook/geelark-phone-webhook-set.md) command to configure the callback URL.
