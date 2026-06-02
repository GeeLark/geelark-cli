# browser update-kernels

Download or update the specified browser kernel to the latest version.

> Available in GeeLark client v4.3.8 and above.

## Key Flags

| Flag | Description |
|------|-------------|
| `--kernel-version <text>` | Browser kernel version, e.g. `"143"` (required) |

## Examples

```bash
# Update kernel
geelark-cli browser update-kernels --kernel-version "143"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | string | Download status: `"complete"` or `"downloading"` |
| `progress` | integer | Download progress percentage |
