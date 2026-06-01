---
name: geelark-shared
version: 1.0.0
description: "Use when first setting up geelark-cli, running config init, verifying authentication, or handling API errors."
---

# geelark-cli Shared Rules

This skill guides you on how to operate GeeLark resources via geelark-cli and important notes.

## Configuration

Run `geelark-cli config init` to set up your API Token on first use.

```bash
# Interactive setup
geelark-cli config init

# Provide token directly
geelark-cli config init --token "your_api_token"
```

You can find your API Token in the GeeLark client Settings page.

## Authentication Verification

After configuration, verify your token is valid:

```bash
geelark-cli auth status
```

- `code: 0` → Authentication successful
- `code: 40003` → Token invalid or expired, re-configure required

## API Response Format

All APIs return a unified JSON envelope:

```json
{
  "traceId": "A25B0025BA886B1EB2679AAAAC599998",
  "code": 0,
  "msg": "success",
  "data": { ... }
}
```

### Envelope Fields (common to all responses)

| Field | Type | Description |
|-------|------|-------------|
| `traceId` | string | Request trace ID; provide this when troubleshooting |
| `code` | integer | `0` means success, non-zero is an error code (see each command's error code table) |
| `msg` | string | Result description; error message on failure |
| `data` | object/null | Business data on success; may be null on failure |

### Success / Failure

- **Success**: `code == 0`, `data` contains business data
- **Failure**: `code != 0`, `msg` contains error description

> The "Response Fields" section in each skill reference only describes the `data` inner structure. The actual response is always wrapped in the envelope above.

### Common Error Codes (Global)

| Code | Description |
|------|-------------|
| 0 | Success |
| 40003 | Token invalid or authentication failed |
| 40004 | Parameter validation failed |
| 40007 | Rate limit exceeded |
| 40011 | Paid plan required |
| 41001 | Insufficient balance |
| 50000 | Internal server error |

Module-specific error codes are listed in each skill reference. Full error codes:
- Cloud Phone: https://open.geelark.com/api/cloud-phone-error-codes
- Browser: https://open.geelark.com/api/browser-error-codes

## Output Format

Control output format with `--format`:

```bash
geelark-cli phone list --format json    # JSON (default)
geelark-cli phone list --format pretty  # Pretty-printed JSON
geelark-cli phone list --format table   # Table
```

## Rate Limits

- Per API: 200 requests/min, 24,000 requests/hour
- Exceeding the limit triggers a 2-hour API ban
- Balance query API: 10 requests/min

## Security Rules

- **Never output tokens** in plain text
- **Confirm user intent** before write/delete operations
- Config file stored at `~/.geelark/config.json`; tokens are masked in display
