# phone app

Application management sub-commands.

## Command Overview

```bash
geelark-cli phone app <command> [flags]
```

| Command | Description |
|---------|-------------|
| `shop-list` | List apps from the GeeLark app store |
| `install` | Install an app on a cloud phone |
| `uninstall` | Uninstall an app from a cloud phone |
| `start` | Start an app on a cloud phone |
| `stop` | Stop an app on a cloud phone |
| `list` | List installed apps on a cloud phone |
| `installable-list` | List apps available for installation |
| `upload` | Upload an app (APK/XAPK) to the team |
| `upload-status` | Query app upload status |
| `batch` | Batch operate apps on cloud phones |
| `team-app` | Team application management (8 sub-commands) |

## shop-list

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number |
| `--page-size <n>` | Page size, max 200 |
| `--key <text>` | Search keyword |
| `--uploaded` | Get uploaded apps only |

```bash
geelark-cli phone app shop-list --page 1 --page-size 10 --key "tiktok"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Page number |
| `pageSize` | integer | Page size |
| `items[]` | array | App data |
| `items[].id` | string | Application ID |
| `items[].appName` | string | Application name |
| `items[].appIcon` | string | Application icon URL |
| `items[].appVersionList[]` | array | Version list |
| `items[].appVersionList[].id` | string | App version ID |
| `items[].appVersionList[].versionCode` | integer | Version code |
| `items[].appVersionList[].versionName` | string | Version name |

## install

| Flag | Description |
|------|-------------|
| `--env-id <text>` | Cloud phone ID (required) |
| `--app-version-id <text>` | App version ID (required) |

```bash
geelark-cli phone app install --env-id "phone_id" --app-version-id "version_id"
```

### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 42003 | App is currently being installed |
| 42004 | A higher version is already installed |
| 42006 | App does not exist |

## uninstall

| Flag | Description |
|------|-------------|
| `--env-id <text>` | Cloud phone ID (required) |
| `--package-name <text>` | Package name (required) |

```bash
geelark-cli phone app uninstall --env-id "phone_id" --package-name "com.zhiliaoapp.musically"
```

### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 42005 | App is not installed |

## start

| Flag | Description |
|------|-------------|
| `--env-id <text>` | Cloud phone ID (required) |
| `--package-name <text>` | Package name (recommended) |
| `--app-version-id <text>` | App version ID (alternative) |

```bash
geelark-cli phone app start --env-id "phone_id" --package-name "com.zhiliaoapp.musically"
```

### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 42003 | App is installing |
| 42005 | App is not installed |

## stop

| Flag | Description |
|------|-------------|
| `--env-id <text>` | Cloud phone ID (required) |
| `--package-name <text>` | Package name (recommended) |
| `--app-version-id <text>` | App version ID (alternative) |

```bash
geelark-cli phone app stop --env-id "phone_id" --package-name "com.zhiliaoapp.musically"
```

### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 42005 | App is not installed |

## list

List installed apps on a cloud phone.

| Flag | Description |
|------|-------------|
| `--env-id <text>` | Cloud phone ID (required) |
| `--page <n>` | Page number |
| `--page-size <n>` | Page size, max 100 |

```bash
geelark-cli phone app list --env-id "phone_id" --page 1 --page-size 10
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Page number |
| `pageSize` | integer | Page size |
| `items[]` | array | App data |
| `items[].appId` | string | Application ID |
| `items[].appName` | string | Application name |
| `items[].appIcon` | string | Application icon URL |
| `items[].appVersionId` | string | Application version ID |
| `items[].packageName` | string | Package name |
| `items[].versionCode` | string | Version code |
| `items[].versionName` | string | Version name |
| `items[].installStatus` | integer | 0=Installing, 1=Installed, 2=Failed, 3=Uninstalling, 4=Uninstalled, 5=Uninstall failed, others=Not installed |
| `items[].installTime` | string | Installation time |

### Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |

## installable-list

List apps available for installation on a cloud phone.

| Flag | Description |
|------|-------------|
| `--env-id <text>` | Cloud phone ID (required) |
| `--page <n>` | Page number |
| `--page-size <n>` | Page size, max 100 |
| `--name <text>` | Search keyword |
| `--uploaded` | Get uploaded apps only |

```bash
geelark-cli phone app installable-list --env-id "phone_id" --name "tiktok"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Page number |
| `pageSize` | integer | Page size |
| `items[]` | array | App data |
| `items[].id` | string | Application ID |
| `items[].appName` | string | Application name |
| `items[].appIcon` | string | Application icon URL |
| `items[].packageName` | string | Package name |
| `items[].installStatus` | integer | 0=Installing, 1=Installed, 2=Failed, 3=Uninstalling, 4=Uninstalled, 5=Uninstall failed, others=Not installed |
| `items[].appVersionInfoList[]` | array | Version list |
| `items[].appVersionInfoList[].id` | string | App version ID |
| `items[].appVersionInfoList[].versionCode` | string | Version code |
| `items[].appVersionInfoList[].versionName` | string | Version name |
| `items[].appVersionInfoList[].installStatus` | integer | 0=Installing, 1=Installed, 2=Failed, 3=Uninstalling, 4=Uninstalled, 5=Uninstall failed, others=Not installed |

### Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |

## upload

Upload an app (APK/XAPK) to the team.

| Flag | Description |
|------|-------------|
| `--url <text>` | Application file URL (required) |
| `--desc <text>` | Description/remark |

```bash
geelark-cli phone app upload --url "https://material.geelark.cn/xxx.apk"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Task ID (use with `upload-status` to query result) |

## upload-status

| Flag | Description |
|------|-------------|
| `--task-id <text>` | Upload task ID (required) |

```bash
geelark-cli phone app upload-status --task-id "task_id"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | 0=Uploading, 1=Success, 2=Failed, 3=Not approved in review |
| `appName` | string | Application name |
| `appIcon` | string | Application icon URL |
| `appId` | string | Application ID |
| `versionId` | string | Application version ID |

### Error Codes

| Code | Description |
|------|-------------|
| 42007 | Task does not exist |

## batch

Batch operate apps on opened cloud phones.

| Flag | Description |
|------|-------------|
| `--action <n>` | Operation: 1=Start, 2=Stop, 3=Restart, 4=Install, 5=Uninstall (required) |
| `--package-name <text>` | Package name (for start/stop/restart/uninstall) |
| `--version-id <text>` | Version ID (for install) |
| `--group-ids <csv>` | Group IDs (default: all groups) |

```bash
# Batch start TikTok
geelark-cli phone app batch --action 1 --package-name "com.zhiliaoapp.musically"

# Batch install
geelark-cli phone app batch --action 4 --version-id "version_id" --group-ids "group1,group2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `items[]` | array | Failed cloud phone entries (only present if some phones fail; successful phones are not listed) |
| `items[].id` | string | Cloud phone ID |
| `items[].errCode` | integer | Error code: 1=App is being installed/uninstalled, 2=App not installed |

## team-app

Team application management sub-commands.

| Sub-command | Description |
|---------|-------------|
| `list` | List team applications |
| `add` | Add an app to team applications |
| `remove` | Remove an app from team applications |
| `set-auth` | Set team app authorization |
| `set-auto-start` | Enable/disable auto-start |
| `set-keep-alive` | Enable/disable keep-alive (Pro plan, Android 12-15) |
| `set-root` | Enable/disable ROOT access |
| `set-auto-install` | Enable/disable auto-install |
| `set-hide` | Set team app hidden status |

### team-app list

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number (required) |
| `--page-size <n>` | Page size, max 200 (required) |

```bash
geelark-cli phone app team-app list --page 1 --page-size 10
```

#### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Page number |
| `pageSize` | integer | Page size |
| `items[]` | array | Team app data |
| `items[].id` | string | Team application ID |
| `items[].appName` | string | Application name |
| `items[].appIcon` | string | Application icon URL |
| `items[].versionId` | string | Version ID |
| `items[].versionCode` | integer | Version code |
| `items[].versionName` | string | Version name |
| `items[].status` | integer | Auto-install: 0=off, 1=on |
| `items[].isUpload` | bool | Whether the app is being uploaded |
| `items[].uploadStatus` | integer | Upload status: 0=Uploading, 1=Success, 2=Failed, 3=Failed review |
| `items[].appAuth` | integer | Authorization: 0=off, 1=on |
| `items[].appRoot` | integer | ROOT: 0=off, 1=on |
| `items[].hideApp` | integer | Hidden: 0=not hidden, 1=hidden |
| `items[].envGroups[]` | array[string] | Allowed group IDs (empty=all groups, "0"=ungrouped) |

### team-app add

| Flag | Description |
|------|-------------|
| `--id <text>` | Application ID (required) |
| `--version-id <text>` | Version ID (required) |
| `--install-group-ids <csv>` | Group IDs (default: all groups; "0" for ungrouped) |

```bash
geelark-cli phone app team-app add --id "app_id" --version-id "version_id"
```

#### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### team-app remove

| Flag | Description |
|------|-------------|
| `--id <text>` | Team application ID (required) |

```bash
geelark-cli phone app team-app remove --id "team_app_id"
```

#### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### team-app set-auth

| Flag | Description |
|------|-------------|
| `--id <text>` | Team application ID (required) |
| `--auth <n>` | Authorization: 0=disable, 1=enable (required) |

```bash
geelark-cli phone app team-app set-auth --id "team_app_id" --auth 1
```

#### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### team-app set-auto-start

| Flag | Description |
|------|-------------|
| `--id <text>` | Team application ID (required) |
| `--opera <n>` | Operation: 0=disable, 1=enable (required) |

```bash
geelark-cli phone app team-app set-auto-start --id "team_app_id" --opera 1
```

#### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### team-app set-keep-alive

Enable/disable keep-alive. Only Pro plan, Android 12-15. Max 1 app can be kept alive.

| Flag | Description |
|------|-------------|
| `--id <text>` | Team application ID (required) |
| `--opera <n>` | Operation: 0=disable, 1=enable (required) |

```bash
geelark-cli phone app team-app set-keep-alive --id "team_app_id" --opera 1
```

#### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

#### Error Codes

| Code | Description |
|------|-------------|
| 44001 | Please upgrade to the Pro plan |

### team-app set-root

| Flag | Description |
|------|-------------|
| `--id <text>` | Team application ID (required) |
| `--opera <n>` | Operation: 0=disable, 1=enable (required) |

```bash
geelark-cli phone app team-app set-root --id "team_app_id" --opera 1
```

#### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### team-app set-auto-install

Enable/disable auto-install. When enabled, the app will be installed after cloud phone starts.

| Flag | Description |
|------|-------------|
| `--id <text>` | Team application ID (required) |
| `--opera <n>` | Operation: 0=disable, 1=enable (required) |

```bash
geelark-cli phone app team-app set-auto-install --id "team_app_id" --opera 1
```

#### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

### team-app set-hide

Set the hidden status of a team app. Only supported on Android 12, 13, and 15.

| Flag | Description |
|------|-------------|
| `--id <text>` | Team application ID (required) |
| `--opera <n>` | Operation: 0=stop hiding, 1=hide (required) |

```bash
geelark-cli phone app team-app set-hide --id "team_app_id" --opera 1
geelark-cli phone app team-app set-hide --id "team_app_id" --opera 0
```

#### Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).
