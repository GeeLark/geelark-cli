# phone analytics

Analytics account management for TikTok, YouTube, Instagram, and Reddit. Requires Pro plan for data queries.

## Command Overview

```bash
geelark-cli phone analytics <command> [flags]
```

| Command | Description |
|---------|-------------|
| `accounts-list` | List analytics accounts |
| `add-accounts` | Batch add analytics accounts |
| `simple-add-account` | Quick add a single analytics account |
| `update-account` | Update an analytics account |
| `delete-account` | Delete an analytics account |
| `data` | Get analytics account data |

### Channel Values

| Value | Platform |
|-------|----------|
| 0 | TikTok |
| 1 | YouTube |
| 2 | Instagram |
| 4 | Reddit |

## accounts-list

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number |
| `--page-size <n>` | Page size, 1-100 |
| `--channel <n>` | Platform filter |
| `--account <text>` | Account name filter |
| `--user-account <text>` | Operator account email |

```bash
geelark-cli phone analytics accounts-list --page 1 --page-size 10
geelark-cli phone analytics accounts-list --channel 0 --account "tk_acc"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `items[]` | array | Account data |
| `items[].id` | string | Account ID |
| `items[].account` | string | Account name |
| `items[].channel` | integer | Platform: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit |
| `items[].remark` | string | Remark |
| `items[].operator` | string | Username of the last operator |
| `items[].created_time` | integer | Creation timestamp (seconds) |
| `items[].updated_time` | integer | Last update timestamp (seconds) |

## add-accounts

Batch add analytics accounts. Max 200 per request.

| Flag | Description |
|------|-------------|
| `--channel <n>` | Platform (required) |
| `--data <json>` | JSON array; each element contains `account` (required, max 64 chars) and `remark` (optional) |

```bash
geelark-cli phone analytics add-accounts --channel 0 --data "[{\"account\":\"acc1\",\"remark\":\"my note\"}]"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `bizCode` | integer | Business status: 0=all successful, 1=account limit exceeded, 2=partially successful with failures |
| `successCount` | integer | Number of successfully added accounts |
| `failCount` | integer | Number of failed additions |
| `repeatCount` | integer | Number of duplicate additions |

## simple-add-account

Quick add a single analytics account.

| Flag | Description |
|------|-------------|
| `--channel <n>` | Platform (required) |
| `--account <text>` | Account name, max 64 chars (required) |
| `--remark <text>` | Remark/note |

```bash
geelark-cli phone analytics simple-add-account --channel 0 --account "myAccount"
```

### Response Fields

Same as [`add-accounts`](#add-accounts).

## update-account

| Flag | Description |
|------|-------------|
| `--id <text>` | Account ID (required) |
| `--account <text>` | New account name, max 64 chars |
| `--channel <n>` | New platform |
| `--remark <text>` | New remark |

```bash
geelark-cli phone analytics update-account --id "id" --account "newName"
```

### Response Fields

Success only (no `data` field). Standard envelope with `code=0` indicates success.

## delete-account

| Flag | Description |
|------|-------------|
| `--channel <n>` | Platform (required) |
| `--account <text>` | Account name (required) |

```bash
geelark-cli phone analytics delete-account --channel 0 --account "myAccount"
```

### Response Fields

Success only (no `data` field). Standard envelope with `code=0` indicates success.

## data

Query analytics account data (play count, follower count, etc.). Requires Pro plan.

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number |
| `--page-size <n>` | Page size, 1-100 |
| `--channel <n>` | Platform filter |
| `--account <text>` | Account name filter |
| `--data-date <n>` | Search date timestamp (seconds) |
| `--created-id <text>` | User ID who added the account |

```bash
geelark-cli phone analytics data --page 1 --page-size 10 --channel 0
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `items[]` | array | Account data |
| `items[].id` | string | Account ID |
| `items[].channel` | integer | Platform |
| `items[].account` | string | Account name |
| `items[].playCount` | integer | Play count (-1 = not yet updated) |
| `items[].followerCount` | integer | Follower count (-1 = not yet updated) |
| `items[].diggCount` | integer | Like/digg count (-1 = not yet updated) |
| `items[].commentCount` | integer | Comment count (-1 = not yet updated) |
| `items[].collectCount` | integer | Collect count (-1 = not yet updated) |
| `items[].shareCount` | integer | Share count (-1 = not yet updated) |
| `items[].dataDate` | integer | Data date timestamp (-1 = not yet updated) |
| `items[].addAccDate` | integer | Account added timestamp (seconds) |
| `items[].remark` | string | Remark |
| `items[].createdId` | string | User ID who added the account |
| `items[].username` | string | Username who added the account |

### Error Codes

| Code | Description |
|------|-------------|
| 43002 | Please upgrade to Pro plan to use this feature |
