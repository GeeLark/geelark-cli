# browser clone

Generate new browsers by cloning an existing one with the same OS and advanced settings.

## Key Flags

| Flag | Description |
|------|-------------|
| `--env-id <text>` | Browser ID to clone (required) |
| `--amount <n>` | Number of clones, 1-100 (default 1) |
| `--group-id <text>` | Target group ID |
| `--clone-name` | Clone the name |
| `--clone-remark` | Clone the remark |
| `--clone-tag` | Clone the tags |
| `--clone-proxy` | Clone the proxy |
| `--clone-cookie` | Clone the cookies |
| `--clone-account` | Clone the account information |

## Examples

```bash
# Clone a browser
geelark-cli browser clone --env-id "browser_id" --amount 2

# Clone with specific settings
geelark-cli browser clone --env-id "browser_id" --amount 1 --clone-name --clone-proxy --clone-cookie
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `ids` | array[string] | Cloned browser IDs |
