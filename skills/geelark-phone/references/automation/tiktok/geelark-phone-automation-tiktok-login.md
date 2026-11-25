# phone automation tiktok-login

TikTok account login.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--account <text>` | Account (max 64 chars, required) |
| `--password <text>` | Password (max 64 chars, required) |
| `--two-fa-key <text>` | 2FA Key |
| `--link-code <text>` | Verification code link (max 999 chars) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Examples

```bash
# Basic login
geelark-cli phone automation tiktok-login --id "557536075321468390" --schedule-at 1741846843 --account "test@gmail.com" --password "123456"

# With 2FA key
geelark-cli phone automation tiktok-login --id "557536075321468390" --schedule-at 1741846843 --account "test@gmail.com" --password "123456" --two-fa-key "2FAKEY"

# With verification code link
geelark-cli phone automation tiktok-login --id "557536075321468390" --schedule-at 1741846843 --account "test@gmail.com" --password "123456" --link-code "https://verification-code-link"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
