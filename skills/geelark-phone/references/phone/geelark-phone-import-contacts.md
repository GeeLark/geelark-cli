# phone import-contacts

Import contacts to a cloud phone. Returns a task ID for querying the result.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--data <json>` | JSON array of contacts (required) |

### Contact Fields (in --data JSON)

| Field | Description |
|-------|-------------|
| `firstName` | First name (at least one of firstName/lastName must be non-empty) |
| `lastName` | Last name |
| `mobile` | Mobile phone number (at least one of mobile/work/fax must be non-empty) |
| `work` | Work phone number |
| `fax` | Fax number |
| `email1` | Email 1 |
| `email2` | Email 2 |

## Examples

```bash
geelark-cli phone import-contacts --id "phone_id" --data "[{\"firstName\":\"Jay\",\"mobile\":\"13288888888\"}]"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Task ID (valid for 1 hour) |
