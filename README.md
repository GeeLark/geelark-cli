# geelark-cli

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.26-blue.svg)](https://go.dev/)
[![npm version](https://img.shields.io/npm/v/geelark-cli.svg)](https://www.npmjs.com/package/geelark-cli)

The official [GeeLark](https://www.geelark.com/) CLI tool — manage cloud phones, browsers, proxies, groups, tags, billing, and more from the command line. Built for humans and AI Agents.

[中文版本请参考这里](#中文版本)

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

- Node.js (`npm`) or Go `v1.26`+ (only required for building from source)

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

**Step 4 — Install Skills (optional)**

Install geelark-cli skills to your AI Agent so it can automatically discover and use geelark-cli commands:

```bash
npx skills@latest add geelark-tech/geelark-cli -g --all -y
```

This registers all geelark-cli skills globally, making them available in Trae, Cursor, Claude Code, Cline, and other supported AI Agents.


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
geelark-cli billing transaction-detail --id "phone_id" --limit 20
geelark-cli billing plan-list
geelark-cli billing plan-info
geelark-cli billing plan-upgrade --profiles-id "497540679501610040" --parallels-num 1 --monthly-rental-num 1
geelark-cli billing plan-renew --days 30
geelark-cli billing buy-time-addon --minutes 2000
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

---

# 中文版本

[GeeLark](https://www.geelark.com/) 官方 CLI 工具 —— 通过命令行管理云手机、浏览器、代理、分组、标签、支付等。为人类和 AI Agent 设计。

## 为什么选择 geelark-cli？

- **全 API 覆盖** — 云手机、浏览器、代理、分组、标签、支付、ADB、Shell、文件、Webhook、OEM、数据助手、应用、自动化、素材中心管理
- **简单快速** — 单一二进制文件，零依赖，Bearer Token 认证
- **AI 友好** — 结构化 JSON 输出，简洁参数，适合自动化场景
- **开源** — MIT 许可证，开箱即用

## 功能概览

| 类别      | 能力 |
|---------| --- |
| 📱 云手机  | 核心：列表、创建、启动、停止、删除、更新、克隆、截图、GPS、重置、Root、短信、转让、联系人 \| 自动化：任务管理、自定义任务、平台自动化 \| 应用：商店、安装/卸载/启动/停止、上传、批量操作、团队应用 \| 文件：临时上传、上传到手机、Keybox \| 素材中心：素材与标签管理 \| ADB：连接信息、启用/禁用 \| Shell：远程命令执行 \| Webhook：回调 URL 管理 \| OEM：品牌定制 \| 数据助手：账号管理与数据查询 |
| 🖥️ 浏览器 | 列表、启动、停止、创建、更新、删除、截图、自动化 |
| 🌐 代理   | 列表、添加、更新、删除、检测代理 |
| 📁 分组   | 列表、创建、更新、删除分组 |
| 🏷️ 标签  | 列表、创建、更新、删除标签 |
| 💰 支付   | 查询账户余额 |

## 安装与快速开始

### 环境要求

- Node.js（`npm`）或 Go `v1.26`+（仅从源码构建时需要）

### 快速开始（人类用户）

#### 安装

**方式一 — 从 npm 安装（推荐）：**

```bash
npm install -g geelark-cli@latest
```

**方式二 — 从源码安装：**

需要 Go `v1.26`+。

```bash
git clone https://github.com/geelark-tech/geelark-cli.git
cd geelark-cli
sudo make install
```

#### 配置与使用

```bash
# 1. 配置 API Token（一次性）
geelark-cli config init --token "your_token_here"

# 或交互式配置：
geelark-cli config init

# 2. 开始使用
geelark-cli phone list --page 1 --page-size 10
```

### 快速开始（AI Agent）

> 以下步骤面向 AI Agent。部分步骤可能需要用户提供 API Token。

**第一步 — 安装**

```bash
npm install -g geelark-cli@latest
```

**第二步 — 配置 API Token**

> 使用用户的 API Token 运行此命令。Token 可从 GeeLark 客户端设置页面获取。

```bash
geelark-cli config init --token "your_token_here"
```

**第三步 — 验证**

```bash
geelark-cli auth status
```

**第四步 — 安装 Skills（可选）**

将 geelark-cli skills 安装到 AI Agent，使其能自动发现和使用 geelark-cli 命令：

```bash
npx skills@latest add geelark-tech/geelark-cli -g --all -y
```

此命令会全局注册所有 geelark-cli skills，使其在 Trae、Cursor、Claude Code、Cline 等支持的 AI Agent 中可用。

## 输出格式

所有命令支持 `--format` 标志控制输出格式：

```bash
# JSON（默认）
geelark-cli phone list --format json

# 格式化 JSON
geelark-cli phone list --format pretty

# 表格格式
geelark-cli phone list --format table
```

## 命令

### `phone` — 云手机管理

所有云手机相关命令都归在 `phone` 下，包括核心操作、自动化、应用、文件、素材中心、ADB、Shell、Webhook、OEM 和数据助手。

#### 核心操作

```bash
# 列出云手机
geelark-cli phone list --page 1 --page-size 10
geelark-cli phone list --ids "id1,id2"
geelark-cli phone list --serial-name "test"
geelark-cli phone list --tags "tag1,tag2"

# 创建云手机
geelark-cli phone create --region "sgp" --mobile-type "Android 12" --data "[{\"profileName\":\"myPhone\",\"proxyInformation\":\"socks5://user:pass@1.2.3.4:1080\"}]"

# 快速创建（简化版）
geelark-cli phone simple-create --region "sgp" --mobile-type "Android 12" --profile-name "myPhone" --proxy-information "socks5://user:pass@1.2.3.4:1080"

# 启动/停止云手机
geelark-cli phone start --ids "id1,id2"
geelark-cli phone stop --ids "id1,id2"

# 删除云手机（需先停止）
geelark-cli phone delete --ids "id1,id2"

# 查询云手机状态
geelark-cli phone status --ids "id1,id2"

# 更新云手机信息
geelark-cli phone update --id "phone_id" --name "new name" --remark "new remark"
geelark-cli phone update --id "phone_id" --tag-ids "tag1,tag2" --group-id "group_id"

# 克隆云手机
geelark-cli phone clone --env-id "phone_id" --amount 2

# 截图
geelark-cli phone screenshot --id "phone_id"
geelark-cli phone screenshot-result --task-id "task_id"

# GPS 管理
geelark-cli phone get-gps --ids "id1,id2"
geelark-cli phone set-gps --data "[{\"id\":\"phone_id\",\"latitude\":1.302,\"longitude\":103.875}]"

# 一键新机（重置身份）
geelark-cli phone new-one --id "phone_id"

# 其他操作
geelark-cli phone set-root --ids "id1,id2" --open
geelark-cli phone send-sms --id "phone_id" --phone-number "+123456" --text "hello"
geelark-cli phone transfer --ids "id1" --transfer-option "name" --account "test@geelark.com"
geelark-cli phone import-contacts --id "phone_id" --data "[{\"firstName\":\"John\",\"mobile\":\"123456\"}]"
```

#### `phone automation` — 自动化任务

```bash
# 任务管理
geelark-cli phone automation task-query --ids "id1,id2"
geelark-cli phone automation task-history --size 10
geelark-cli phone automation task-cancel --ids "id1,id2"
geelark-cli phone automation task-restart --ids "id1,id2"

# 自定义任务
geelark-cli phone automation add-custom-task --id "phone_id" --schedule-at 1741846843 --flow-id "flow_id" --param-map "{\"key\":\"value\"}"

# 平台自动化（TikTok、Facebook、Instagram、YouTube 等）
geelark-cli phone automation tiktok-login --id "phone_id" --schedule-at 1741846843 --account "user" --password "pass"
geelark-cli phone automation facebook-maintenance --id "557536075321468390" --schedule-at 1741846843 --browse-posts-num 10 --keyword "k1,k2"
geelark-cli phone automation instagram-login --id "phone_id" --schedule-at 1741846843 --account "user" --password "pass"
```

#### `phone app` — 应用管理

```bash
# 应用商店
geelark-cli phone app shop-list --page 1 --page-size 10 --key "tiktok"

# 安装/卸载/启动/停止
geelark-cli phone app install --env-id "phone_id" --app-version-id "version_id"
geelark-cli phone app uninstall --env-id "phone_id" --package-name "com.example.app"
geelark-cli phone app start --env-id "phone_id" --package-name "com.example.app"
geelark-cli phone app stop --env-id "phone_id" --package-name "com.example.app"

# 查看已安装/可安装应用
geelark-cli phone app list --env-id "phone_id" --page 1
geelark-cli phone app installable-list --env-id "phone_id" --name "tiktok"

# 上传与状态查询
geelark-cli phone app upload --url "https://material.geelark.cn/xxx.apk"
geelark-cli phone app upload-status --task-id "task_id"

# 批量操作（1=启动, 2=停止, 3=重启, 4=安装, 5=卸载）
geelark-cli phone app batch --action 1 --package-name "com.example.app"
geelark-cli phone app batch --action 4 --version-id "version_id" --group-ids "group1,group2"

# 团队应用管理
geelark-cli phone app team-app list --page 1
geelark-cli phone app team-app add --id "app_id" --version-id "version_id"
geelark-cli phone app team-app remove --id "team_app_id"
geelark-cli phone app team-app set-auth --id "team_app_id" --auth 1
geelark-cli phone app team-app set-auto-start --id "team_app_id" --opera 1
geelark-cli phone app team-app set-keep-alive --id "team_app_id" --opera 1
geelark-cli phone app team-app set-root --id "team_app_id" --opera 1
geelark-cli phone app team-app set-auto-install --id "team_app_id" --status 1
```

#### `phone file` — 文件管理

```bash
# 上传临时文件到 GeeLark（有效期 3 天）
geelark-cli phone file upload-temp --file ./video.mp4

# 上传文件到云手机（通过 URL 或本地文件）
geelark-cli phone file upload-to-phone --id "phone_id" --url "https://material.geelark.cn/xxx.mp4"
geelark-cli phone file upload-to-phone --id "phone_id" --file ./local-video.mp4

# 查询上传状态
geelark-cli phone file upload-status --task-id "task_id"

# Keybox 上传
geelark-cli phone file keybox-upload --id "phone_id" --file ./keybox.xml
geelark-cli phone file keybox-result --task-id "task_id"
```

#### `phone library` — 素材与标签管理

```bash
# 上传并创建素材（合并获取上传地址 + PUT + 创建）
geelark-cli phone library material-create --file ./photo.jpg
geelark-cli phone library material-create --file ./video.mp4 --tags-id "tag1,tag2"

# 搜索/删除素材
geelark-cli phone library material-search --page 1 --page-size 10
geelark-cli phone library material-delete --ids "id1,id2"

# 标签管理
geelark-cli phone library tag-create --name "myTag"
geelark-cli phone library tag-search --page 1
geelark-cli phone library tag-set --materials-id "id1,id2" --tags-id "tag1,tag2"
geelark-cli phone library tag-delete --ids "id1,id2"
```

#### `phone adb` — ADB 管理

```bash
# 获取 ADB 连接信息
geelark-cli phone adb get-info --ids "id1,id2"

# 启用/禁用 ADB
geelark-cli phone adb set-status --ids "id1,id2" --open
geelark-cli phone adb set-status --ids "id1" --open=false
```

#### `phone shell` — Shell 执行

```bash
# 在云手机上执行 Shell 命令
geelark-cli phone shell exec --id "phone_id" --cmd "pm list packages"
geelark-cli phone shell exec --id "phone_id" --cmd "ls /sdcard/Download"
```

#### `phone webhook` — Webhook 管理

```bash
# 获取当前 Webhook URL
geelark-cli phone webhook get

# 设置 Webhook URL
geelark-cli phone webhook set --url "https://example.com/callback"
```

#### `phone oem` — OEM / 白标

```bash
# 自定义品牌设置
geelark-cli phone oem customization --title "MyBrand" --logo "https://example.com/logo.png"
geelark-cli phone oem customization --hide-header --mirror-url "https://www.abcd.com/mirror/url"
geelark-cli phone oem customization --toolbar "[{\"toolBar\":\"networkQuality\",\"visible\":false}]"
```

#### `phone analytics` — 数据助手

```bash
# 列出数据助手账号
geelark-cli phone analytics accounts-list --page 1 --page-size 10

# 添加/更新/删除账号
geelark-cli phone analytics simple-add-account --channel 0 --account "myAccount"
geelark-cli phone analytics update-account --id "account_id" --account "newName"
geelark-cli phone analytics delete-account --channel 0 --account "myAccount"

# 查询数据
geelark-cli phone analytics data --page 1 --channel 0
```

### `browser` — 浏览器管理

```bash
# 列出浏览器
geelark-cli browser list --page 1 --page-size 10

# 创建浏览器
geelark-cli browser create --data "{\"serialName\":\"myBrowser\",\"browserOs\":1}"
geelark-cli browser create --data "{\"serialName\":\"test\",\"browserOs\":2,\"accountPlatform\":\"https://www.tiktok.com/\",\"accountUsername\":\"user\",\"accountPassword\":\"pass\"}"

# 编辑浏览器
geelark-cli browser edit --data "{\"id\":\"browser_id\",\"serialName\":\"newName\"}"

# 启动/停止浏览器
geelark-cli browser start --id "id1"
geelark-cli browser stop --id "id1"

# 浏览器自动化
geelark-cli browser automation add-custom-task --eid "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885" --param-map "{\"Title\":\"video\",\"Desc\":\"this is video\"}"
```

### `proxy` — 代理管理

```bash
# 列出代理
geelark-cli proxy list --page 1 --page-size 10

# 添加代理
geelark-cli proxy add --data "[{\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000,\"username\":\"admin\",\"password\":\"admin\"}]"

# 更新代理
geelark-cli proxy update --data "[{\"id\":\"proxy_id\",\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000}]"

# 删除代理
geelark-cli proxy delete --ids "id1,id2"

# 检测代理
geelark-cli proxy check --type socks5 --server 185.162.130.86 --port 10000 --channel IP2Location
```

### `group` — 分组管理

```bash
geelark-cli group list --page 1 --page-size 10
geelark-cli group create --data "[{\"name\":\"myGroup\",\"remark\":\"description\"}]"
geelark-cli group update --data "[{\"id\":\"group_id\",\"name\":\"newName\"}]"
geelark-cli group delete --ids "id1,id2"
```

### `tag` — 标签管理

```bash
geelark-cli tag list --page 1 --page-size 10
geelark-cli tag create --data "[{\"name\":\"myTag\",\"color\":\"red\"}]"
geelark-cli tag update --data "[{\"id\":\"tag_id\",\"name\":\"newName\",\"color\":\"blue\"}]"
geelark-cli tag delete --ids "id1,id2"
```

### `billing` — 支付

```bash
geelark-cli billing balance
geelark-cli billing transaction-detail --id "phone_id" --limit 20
geelark-cli billing plan-list
geelark-cli billing plan-info
geelark-cli billing plan-upgrade --profiles-id "497540679501610040" --parallels-num 1 --monthly-rental-num 1
geelark-cli billing plan-renew --days 30
geelark-cli billing buy-time-addon --minutes 2000
```

## 项目结构

```
geelark-cli/
├── main.go                    # 入口
├── go.mod                     # Go 模块文件
├── Makefile                   # 构建自动化
├── README.md                  # 本文件
├── cmd/                       # 命令实现
│   ├── root.go               # 根命令与全局标志
│   ├── config/               # 配置管理（init、show）
│   ├── phone/                # 云手机及所有子模块
│   │   ├── phone.go          #   云手机核心命令
│   │   ├── automation.go     #   自动化任务与平台操作
│   │   ├── app.go            #   应用与团队应用管理
│   │   ├── file.go           #   文件上传与 Keybox
│   │   ├── library.go        #   素材与标签管理
│   │   ├── shell.go          #   Shell 执行
│   │   ├── adb.go            #   ADB 管理
│   │   ├── webhook.go        #   Webhook 管理
│   │   ├── oem.go            #   OEM定制
│   │   └── analytics.go      #   数据助手账号管理
│   ├── browser/              # 浏览器管理
│   │   ├── browser.go        #   浏览器核心命令
│   │   └── automation.go     #   浏览器自动化
│   ├── proxy/                # 代理管理
│   ├── group/                # 分组管理
│   ├── tag/                  # 标签管理
│   └── billing/              # 支付与订阅
└── internal/                  # 内部包
    ├── build/                # 构建版本信息
    ├── client/               # API 客户端（HTTP + Bearer 认证 + 文件上传）
    ├── config/               # 配置文件管理
    └── output/               # 输出格式化（json/pretty/table）
```

## API 限流

- 每个接口限流：200 次/分钟，24,000 次/小时
- 超出限制后，API 将被限制 2 小时
- 余额查询接口：10 次/分钟

## 错误码

### 云手机

| 错误码 | 描述 |
| --- | --- |
| 0 | 成功 |
| 40004 | 参数校验失败 |
| 40007 | 请求频率超限 |
| 40011 | 仅付费用户可用 |
| 41001 | 余额不足 |
| 42001 | 云手机不存在 |
| 42002 | 云手机未在运行状态 |
| 43005 | 云手机正在执行任务 |
| 44002 | 配置数量已达套餐上限 |
| 48002 | 任务不存在 |
| 50000 | 服务器内部错误 |

完整错误码：https://open.geelark.cn/api/cloud-phone-error-codes

### 浏览器

| 错误码 | 描述 |
| --- | --- |
| 0 | 成功 |
| 40004 | 请求参数无效 |
| 40007 | 请求频率超限 |
| 40011 | 仅付费用户可用 |
| 44002 | 环境数量已达套餐上限 |
| 50000 | 服务器内部错误 |
| 90001 | 用户未登录 |
| 90007 | 请求头缺少 Authorization |

完整错误码：https://open.geelark.cn/api/browser-error-codes

## 贡献

1. Fork 本仓库
2. 创建功能分支（`git checkout -b feature/amazing-feature`）
3. 提交更改（`git commit -m 'Add amazing feature'`）
4. 推送到分支（`git push origin feature/amazing-feature`）
5. 发起 Pull Request

## 许可证

本项目基于 MIT 许可证 —— 详见 [LICENSE](LICENSE) 文件。

## 链接

- **GeeLark 官网**：https://www.geelark.com
- **API 文档**：https://open.geelark.cn/api
- **GitHub Issues**：https://github.com/geelark-tech/geelark-cli/issues
