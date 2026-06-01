# phone automation task-flow-export

Export a custom task flow.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Custom task flow ID (required) |

## Example

```bash
geelark-cli phone automation task-flow-export --id "612345671223083526"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `gal` | string | Custom task flow data (JSON string) |

### Error Codes

| Code | Description |
|------|-------------|
| 48002 | Custom task flow not found |
