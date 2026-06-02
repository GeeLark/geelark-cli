# browser get-kernels

Query the list of available browser kernels.

> Available in GeeLark client v4.3.8 and above.

## Examples

```bash
# List available kernels
geelark-cli browser get-kernels
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `data[]` | array | Kernel list |
| `data[].kernel` | string | Kernel version |
| `data[].isDownloaded` | bool | Whether the kernel has been downloaded |
