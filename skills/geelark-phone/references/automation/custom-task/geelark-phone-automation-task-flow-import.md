# phone automation task-flow-import

Import or update a custom task flow. If `--id` is provided, the existing flow is updated; otherwise a new one is created.


## Flags

| Flag | Description |
|------|-------------|
| `--gal <text>` | Custom task flow data (required) |
| `--id <text>` | Custom task flow ID (for update; omit to create new) |

## Example

```bash
geelark-cli phone automation task-flow-import --gal '{"content":{...}}'
geelark-cli phone automation task-flow-import --id "612345671223083526" --gal '{"content":{...}}'
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Custom task flow ID |

### Error Codes

| Code | Description |
|------|-------------|
| 48002 | Custom task flow not found |
