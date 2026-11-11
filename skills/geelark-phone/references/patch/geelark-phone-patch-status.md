# phone patch status

Query the install status of an Xposed patch. Supports Android 12/13/14/15/16.

- `--type 1`: query by patch ID (Android 12/13/15 only, requires `--patch-id` and `--id`).
- `--type 2`: query by task ID (all supported Android versions, requires `--task-id`).

## Key Flags

| Flag | Description |
|------|-------------|
| `--type <n>` | Query type: 1=by patchId (Android 12/13/15), 2=by taskId (all) (required) |
| `--patch-id <text>` | Patch ID, required when type=1 |
| `--id <text>` | Cloud phone ID, required when type=1 |
| `--task-id <text>` | Task ID returned by `patch install`, required when type=2 |

## Examples

```bash
# Query by task ID (all Android versions)
geelark-cli phone patch status --type 2 --task-id "1830906144634757120"

# Query by patch ID (Android 12/13/15)
geelark-cli phone patch status --type 1 --patch-id "sdktest" --id "532930887177274368"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | Install status: 0=installing, 1=success, 2=upload failed, 3=execution failed or success cannot be confirmed |

## Error Codes

| Code | Description |
|------|-------------|
| 42007 | Corresponding task does not exist |
