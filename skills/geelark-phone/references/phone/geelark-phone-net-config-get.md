# phone net-config-get

Get cloud phone network settings including access blacklist.

## Examples

```bash
geelark-cli phone net-config-get
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `blackList` | array[string] | Blacklisted domains |
