# geelark-cli

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.26-blue.svg)](https://go.dev/)
[![npm version](https://img.shields.io/npm/v/geelark-cli.svg)](https://www.npmjs.com/package/geelark-cli)

The official [GeeLark](https://www.geelark.com/) CLI tool — manage cloud phones, browsers, proxies, groups, tags, billing, and more from the command line. Built for humans and AI Agents.

## Why geelark-cli?

- **Full API Coverage** — Cloud phone, browser, proxy, group, tag, billing, ADB, shell, file, webhook, OEM, analytics, app, automation, and library management
- **Simple & Fast** — One binary, zero dependencies, Bearer token auth
- **AI-Friendly** — Structured JSON output, concise parameters, perfect for automation
- **Open Source** — MIT license, ready to use

## Features

| Category | Capabilities |
| --- | --- |
| 📱 Phone | Core: list, create, start, stop, delete, update, clone, screenshot, GPS, reset, root, SMS, transfer, contacts \| Automation: task management, custom tasks, platform automation \| App: shop, install/uninstall/start/stop, upload, batch, team app \| File: temp upload, upload to phone, keybox \| Library: material & tag management \| ADB: connection info, enable/disable \| Shell: remote command execution \| Webhook: callback URL management \| OEM: brand customization \| Analytics: account management & data query |
| 🖥️ Browser | List, start, stop, create, update, delete, screenshot, automation |
| 🌐 Proxy | List, add, update, delete, check/detect proxies |
| 📁 Group | List, create, update, delete groups |
| 🏷️ Tag | List, create, update, delete tags |
| 💰 Billing | Query account balance |

## Installation & Quick Start

### Requirements

- Node.js (`npm`/`npx`) or Go `v1.26`+ (only required for building from source)

### Quick Start (Human Users)

#### Install

**Option 1 — Install from npm (recommended):**

```bash
npm install -g geelark-cli@latest
```

**Option 2 — Install from source:**

Requires Go `v1.26`+.

```bash
git clone https://github.com/geelark-tech/geelark-cli.git
cd geelark-cli
sudo make install
```

#### Configure & Use

```bash
# 1. Configure your API token (one-time setup)
geelark-cli config init --token "your_token_here"

# Or interactively:
geelark-cli config init

# 2. Start using
geelark-cli phone list --page 1 --page-size 10
```

### Quick Start (AI Agent)

> The following steps are for AI Agents. Some steps may require the user to provide their API token.

**Step 1 — Install**

```bash
npm install -g geelark-cli@latest
```

**Step 2 — Configure API token**

> Run this command with the user's API token. The token can be obtained from the GeeLark client settings page.

```bash
geelark-cli config init --token "your_token_here"
```

**Step 3 — Verify**

```bash
geelark-cli auth status
```


## Output Format

All commands support `--format` flag to control output:

```bash
# JSON (default)
geelark-cli phone list --format json

# Pretty-printed JSON
geelark-cli phone list --format pretty

# Table format
geelark-cli phone list --format table
```

## Commands

### `phone` — Cloud Phone Management

All cloud phone related commands are grouped under `phone`, including core operations, automation, app, file, library, ADB, shell, webhook, OEM, and analytics.

#### Core Operations

```bash
# List cloud phones
geelark-cli phone list --page 1 --page-size 10
geelark-cli phone list --ids "id1,id2"
geelark-cli phone list --serial-name "test"
geelark-cli phone list --tags "tag1,tag2"

# Create cloud phones
geelark-cli phone create --region "sgp" --mobile-type "Android 12" --data "[{\"profileName\":\"myPhone\",\"proxyInformation\":\"socks5://user:pass@1.2.3.4:1080\"}]"

# Quick create (simplified)
geelark-cli phone simple-create --region "sgp" --mobile-type "Android 12" --profile-name "myPhone" --proxy-information "socks5://user:pass@1.2.3.4:1080"

# Start/Stop cloud phones
geelark-cli phone start --ids "id1,id2"
geelark-cli phone stop --ids "id1,id2"

# Delete cloud phones (must be stopped first)
geelark-cli phone delete --ids "id1,id2"

# Query cloud phone status
geelark-cli phone status --ids "id1,id2"

# Update cloud phone info
geelark-cli phone update --id "phone_id" --name "new name" --remark "new remark"
geelark-cli phone update --id "phone_id" --tag-ids "tag1,tag2" --group-id "group_id"

# Clone a cloud phone
geelark-cli phone clone --env-id "phone_id" --amount 2

# Take a screenshot
geelark-cli phone screenshot --id "phone_id"
geelark-cli phone screenshot-result --task-id "task_id"

# GPS management
geelark-cli phone get-gps --ids "id1,id2"
geelark-cli phone set-gps --data "[{\"id\":\"phone_id\",\"latitude\":1.302,\"longitude\":103.875}]"

# One-click new machine (reset identity)
geelark-cli phone new-one --id "phone_id"

# Other operations
geelark-cli phone set-root --ids "id1,id2" --open
geelark-cli phone send-sms --id "phone_id" --phone-number "+123456" --text "hello"
geelark-cli phone transfer --ids "id1" --transfer-option "name" --account "test@geelark.com"
geelark-cli phone import-contacts --id "phone_id" --data "[{\"firstName\":\"John\",\"mobile\":\"123456\"}]"
```

#### `phone automation` — Automation Tasks

```bash
# Task management
geelark-cli phone automation task-query --ids "id1,id2"
geelark-cli phone automation task-history --size 10
geelark-cli phone automation task-cancel --ids "id1,id2"
geelark-cli phone automation task-restart --ids "id1,id2"

# Custom tasks
geelark-cli phone automation add-custom-task --id "phone_id" --schedule-at 1741846843 --flow-id "flow_id" --param-map "{\"key\":\"value\"}"

# Platform automation (TikTok, Facebook, Instagram, YouTube, etc.)
geelark-cli phone automation tiktok-login --id "phone_id" --schedule-at 1741846843 --account "user" --password "pass"
geelark-cli phone automation facebook-maintenance --id "557536075321468390" --schedule-at 1741846843 --browse-posts-num 10 --keyword "k1,k2"
geelark-cli phone automation instagram-login --id "phone_id" --schedule-at 1741846843 --account "user" --password "pass"
```

#### `phone app` — Application Management

```bash
# App store
geelark-cli phone app shop-list --page 1 --page-size 10 --key "tiktok"

# Install/Uninstall/Start/Stop
geelark-cli phone app install --env-id "phone_id" --app-version-id "version_id"
geelark-cli phone app uninstall --env-id "phone_id" --package-name "com.example.app"
geelark-cli phone app start --env-id "phone_id" --package-name "com.example.app"
geelark-cli phone app stop --env-id "phone_id" --package-name "com.example.app"

# List installed / installable apps
geelark-cli phone app list --env-id "phone_id" --page 1
geelark-cli phone app installable-list --env-id "phone_id" --name "tiktok"

# Upload & status
geelark-cli phone app upload --url "https://material.geelark.cn/xxx.apk"
geelark-cli phone app upload-status --task-id "task_id"

# Batch operations (1=Start, 2=Stop, 3=Restart, 4=Install, 5=Uninstall)
geelark-cli phone app batch --action 1 --package-name "com.example.app"
geelark-cli phone app batch --action 4 --version-id "version_id" --group-ids "group1,group2"

# Team app management
geelark-cli phone app team-app list --page 1
geelark-cli phone app team-app add --id "app_id" --version-id "version_id"
geelark-cli phone app team-app remove --id "team_app_id"
geelark-cli phone app team-app set-auth --id "team_app_id" --auth 1
geelark-cli phone app team-app set-auto-start --id "team_app_id" --opera 1
geelark-cli phone app team-app set-keep-alive --id "team_app_id" --opera 1
geelark-cli phone app team-app set-root --id "team_app_id" --opera 1
geelark-cli phone app team-app set-auto-install --id "team_app_id" --status 1
```

#### `phone file` — File Management

```bash
# Upload temp file to GeeLark (valid 3 days)
geelark-cli phone file upload-temp --file ./video.mp4

# Upload file to cloud phone (by URL or local file)
geelark-cli phone file upload-to-phone --id "phone_id" --url "https://material.geelark.cn/xxx.mp4"
geelark-cli phone file upload-to-phone --id "phone_id" --file ./local-video.mp4

# Query upload status
geelark-cli phone file upload-status --task-id "task_id"

# Keybox upload
geelark-cli phone file keybox-upload --id "phone_id" --file ./keybox.xml
geelark-cli phone file keybox-result --task-id "task_id"
```

#### `phone library` — Material & Tag Management

```bash
# Upload and create material (combines getUploadUrl + PUT + create)
geelark-cli phone library material-create --file ./photo.jpg
geelark-cli phone library material-create --file ./video.mp4 --tags-id "tag1,tag2"

# Search / delete materials
geelark-cli phone library material-search --page 1 --page-size 10
geelark-cli phone library material-delete --ids "id1,id2"

# Tag management
geelark-cli phone library tag-create --name "myTag"
geelark-cli phone library tag-search --page 1
geelark-cli phone library tag-set --materials-id "id1,id2" --tags-id "tag1,tag2"
geelark-cli phone library tag-delete --ids "id1,id2"
```

#### `phone adb` — ADB Management

```bash
# Get ADB connection info
geelark-cli phone adb get-info --ids "id1,id2"

# Enable/disable ADB
geelark-cli phone adb set-status --ids "id1,id2" --open
geelark-cli phone adb set-status --ids "id1" --open=false
```

#### `phone shell` — Shell Execution

```bash
# Execute a shell command on a cloud phone
geelark-cli phone shell exec --id "phone_id" --cmd "pm list packages"
geelark-cli phone shell exec --id "phone_id" --cmd "ls /sdcard/Download"
```

#### `phone webhook` — Webhook Management

```bash
# Get current webhook URL
geelark-cli phone webhook get

# Set webhook URL
geelark-cli phone webhook set --url "https://example.com/callback"
```

#### `phone oem` — OEM / White Label

```bash
# Customize brand settings
geelark-cli phone oem customization --title "MyBrand" --logo "https://example.com/logo.png"
geelark-cli phone oem customization --hide-header --mirror-url "https://www.abcd.com/mirror/url"
geelark-cli phone oem customization --toolbar "[{\"toolBar\":\"networkQuality\",\"visible\":false}]"
```

#### `phone analytics` — Analytics Account Management

```bash
# List analytics accounts
geelark-cli phone analytics accounts-list --page 1 --page-size 10

# Add / update / delete accounts
geelark-cli phone analytics simple-add-account --channel 0 --account "myAccount"
geelark-cli phone analytics update-account --id "account_id" --account "newName"
geelark-cli phone analytics delete-account --channel 0 --account "myAccount"

# Query analytics data
geelark-cli phone analytics data --page 1 --channel 0
```

### `browser` — Browser Management

```bash
# List browsers
geelark-cli browser list --page 1 --page-size 10

# Create browser
geelark-cli browser create --data "{\"serialName\":\"myBrowser\",\"browserOs\":1}"
geelark-cli browser create --data "{\"serialName\":\"test\",\"browserOs\":2,\"accountPlatform\":\"https://www.tiktok.com/\",\"accountUsername\":\"user\",\"accountPassword\":\"pass\"}"

# Edit browser
geelark-cli browser edit --data "{\"id\":\"browser_id\",\"serialName\":\"newName\"}"

# Start/Stop browsers
geelark-cli browser start --id "id1"
geelark-cli browser stop --id "id1"

# Browser automation
geelark-cli browser automation add-custom-task --eid "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885" --param-map "{\"Title\":\"video\",\"Desc\":\"this is video\"}"
```

### `proxy` — Proxy Management

```bash
# List proxies
geelark-cli proxy list --page 1 --page-size 10

# Add proxies
geelark-cli proxy add --data "[{\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000,\"username\":\"admin\",\"password\":\"admin\"}]"

# Update proxies
geelark-cli proxy update --data "[{\"id\":\"proxy_id\",\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000}]"

# Delete proxies
geelark-cli proxy delete --ids "id1,id2"

# Check/detect a proxy
geelark-cli proxy check --type socks5 --server 185.162.130.86 --port 10000 --channel IP2Location
```

### `group` — Group Management

```bash
geelark-cli group list --page 1 --page-size 10
geelark-cli group create --data "[{\"name\":\"myGroup\",\"remark\":\"description\"}]"
geelark-cli group update --data "[{\"id\":\"group_id\",\"name\":\"newName\"}]"
geelark-cli group delete --ids "id1,id2"
```

### `tag` — Tag Management

```bash
geelark-cli tag list --page 1 --page-size 10
geelark-cli tag create --data "[{\"name\":\"myTag\",\"color\":\"red\"}]"
geelark-cli tag update --data "[{\"id\":\"tag_id\",\"name\":\"newName\",\"color\":\"blue\"}]"
geelark-cli tag delete --ids "id1,id2"
```

### `billing` — Billing

```bash
geelark-cli billing balance
```

## Project Structure

```
geelark-cli/
├── main.go                    # Entry point
├── go.mod                     # Go module file
├── Makefile                   # Build automation
├── README.md                  # This file
├── cmd/                       # Command implementations
│   ├── root.go               # Root command & global flags
│   ├── config/               # Config management (init, show)
│   ├── phone/                # Cloud phone & all sub-modules
│   │   ├── phone.go          #   Phone core commands
│   │   ├── automation.go     #   Automation tasks & platform ops
│   │   ├── app.go            #   App & team app management
│   │   ├── file.go           #   File upload & keybox
│   │   ├── library.go        #   Material & tag management
│   │   ├── shell.go          #   Shell execution
│   │   ├── adb.go            #   ADB management
│   │   ├── webhook.go        #   Webhook management
│   │   ├── oem.go            #   OEM/White Label customization
│   │   └── analytics.go      #   Analytics account management
│   ├── browser/              # Browser management
│   │   ├── browser.go        #   Browser core commands
│   │   └── automation.go     #   Browser automation
│   ├── proxy/                # Proxy management
│   ├── group/                # Group management
│   ├── tag/                  # Tag management
│   └── billing/              # Billing & subscription
└── internal/                  # Internal packages
    ├── build/                # Build version info
    ├── client/               # API client (HTTP + Bearer auth + file upload)
    ├── config/               # Config file management
    └── output/               # Output formatting (json/pretty/table)
```

## API Rate Limits

- Per API rate limit: 200 requests per minute, 24,000 requests per hour
- After exceeding the limit, the API will be restricted for 2 hours
- Balance API: 10 requests per minute

## Error Codes

### Cloud Phone

| Code | Description |
| --- | --- |
| 0 | Success |
| 40004 | Parameter validation failed |
| 40007 | Request rate limited |
| 40011 | Available to paid users only |
| 41001 | Insufficient balance |
| 42001 | Cloud phone not found |
| 42002 | Cloud phone is not in running state |
| 43005 | Cloud phone is executing a task |
| 44002 | Profiles count has reached the plan limit |
| 48002 | Task does not exist |
| 50000 | Internal server error |

Full error codes: https://open.geelark.com/api/cloud-phone-error-codes

### Browser

| Code | Description |
| --- | --- |
| 0 | Success |
| 40004 | Invalid request parameters |
| 40007 | Request rate limit exceeded |
| 40011 | This API is available to paid users only |
| 44002 | Environment count has reached the plan limit |
| 50000 | Internal server error |
| 90001 | User not logged in |
| 90007 | Missing 'Authorization' in the header |

Full error codes: https://open.geelark.com/api/browser-error-codes

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Links

- **GeeLark Website**: https://www.geelark.com
- **API Documentation**: https://open.geelark.com/api
- **GitHub Issues**: https://github.com/geelark-tech/geelark-cli/issues
