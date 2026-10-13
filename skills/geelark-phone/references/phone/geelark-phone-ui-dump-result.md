# phone ui-dump-result

Query the status of a cloud phone UI Dump task. The result can be retrieved within 30 minutes after requesting [`ui-dump`](geelark-phone-ui-dump.md); after expiry, retrieval fails.

## Key Flags

| Flag | Description |
|------|-------------|
| `--task-id <text>` | UI Dump task ID (required) |

## Examples

```bash
geelark-cli phone ui-dump-result --task-id "task_id"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | 0=Acquisition failed, 1=In progress, 2=Execution succeeded, 3=Execution failed |
| `downloadLink` | string | UI layout xml file download link (present when succeeded) |

## Error Codes

For error codes, see the shared error code table in [geelark-phone SKILL](../../SKILL.md#error-codes).
