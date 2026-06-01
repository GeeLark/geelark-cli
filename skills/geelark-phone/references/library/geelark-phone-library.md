# phone library

Library (material) management sub-commands. Upload, search, delete materials and manage tags.

## Command Overview

```bash
geelark-cli phone library <command> [flags]
```

| Command | Description |
|---------|-------------|
| `material-create` | Upload a file and create a material |
| `material-search` | Search materials |
| `material-delete` | Delete materials |
| `tag-create` | Create a material tag |
| `tag-search` | Search material tags |
| `tag-delete` | Delete material tags |
| `tag-set` | Set tags on materials |

## material-create

Upload a local file and create a material in the Library. Combines getUploadUrl + upload + create in one step.

| Flag | Description |
|------|-------------|
| `--file <path>` | Local file path (required) |
| `--file-name <text>` | Material name (defaults to file name, max 200 chars) |
| `--tags-id <csv>` | Tag IDs |

```bash
geelark-cli phone library material-create --file ./photo.jpg
geelark-cli phone library material-create --file ./video.mp4 --tags-id "tag1,tag2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Material ID |
| `failDetails[]` | array | Tag binding failure details (only present if some tags fail) |
| `failDetails[].id` | string | Tag ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

### Error Codes

| Code | Description |
|------|-------------|
| 60001 | Material library has reached maximum capacity |
| 60003 | Illegal URL, please upload temporary files first |
| 60004 | File format not supported |
| 43022 | Tag not found |
| 40005 | Resource does not exist, check URL availability |

## material-search

Search materials in the Library.

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number (default 1) |
| `--page-size <n>` | Page size, max 200 (default 50) |
| `--file-name <text>` | Search by file name |
| `--tags-id <csv>` | Filter by tag IDs |
| `--source <n>` | Source: 0=upload, 1=AI Edit, 2=Baidu Cloud, 3=GhostCut, 4=GoogleDrive, 5=Image to video |
| `--file-type <csv>` | File types: 1=image, 2=video |
| `--ids <csv>` | Material IDs, max 100 |

```bash
geelark-cli phone library material-search --page 1 --page-size 50
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `list[]` | array | Material data |
| `list[].id` | string | Material ID |
| `list[].createdTime` | integer | Creation timestamp (seconds) |
| `list[].fileName` | string | Material name |
| `list[].fileSize` | integer | File size in bytes |
| `list[].fileUrl` | string | Material URL |
| `list[].fileType` | integer | File type: 1=image, 2=video |
| `list[].width` | integer | Width (px) |
| `list[].height` | integer | Height (px) |
| `list[].source` | integer | Source: 0=upload, 1=AI Edit, 2=Baidu Cloud, 3=GhostCut, 4=GoogleDrive, 5=Image to video |
| `list[].tags[]` | array | Tag data |
| `list[].tags[].id` | string | Tag ID |
| `list[].tags[].name` | string | Tag name |
| `list[].tags[].color` | integer | Tag color: 0=White, 1=Red, 2=Blue, 3=Green, 4=Yellow, 5=Purple |
| `list[].userName` | string | Upload user name |

## material-delete

Delete materials from the Library.

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Material IDs (required) |

```bash
geelark-cli phone library material-delete --ids "id1,id2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total delete requests |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `failDetails[]` | array | Failure details |
| `failDetails[].id` | string | Material ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

### Error Codes

| Code | Description |
|------|-------------|
| 60005 | Material not found |

## tag-create

Create a material tag.

| Flag | Description |
|------|-------------|
| `--name <text>` | Tag name, max 30 chars (required) |
| `--color <n>` | Color: 0=White, 1=Red, 2=Blue, 3=Green, 4=Yellow, 5=Purple |

```bash
geelark-cli phone library tag-create --name "my-tag" --color 1
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Tag ID |

### Error Codes

| Code | Description |
|------|-------------|
| 60002 | Tag name already exists |

## tag-search

Search material tags.

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number |
| `--page-size <n>` | Page size, max 200 |
| `--name <text>` | Search by tag name |

```bash
geelark-cli phone library tag-search --page 1 --page-size 50
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `list[]` | array | Tag data |
| `list[].id` | string | Tag ID |
| `list[].name` | string | Tag name |
| `list[].color` | integer | Tag color: 0=White, 1=Red, 2=Blue, 3=Green, 4=Yellow, 5=Purple |

## tag-delete

Delete material tags.

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Tag IDs (required) |

```bash
geelark-cli phone library tag-delete --ids "id1,id2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `totalAmount` | integer | Total delete requests |
| `successAmount` | integer | Successful count |
| `failAmount` | integer | Failed count |
| `failDetails[]` | array | Failure details |
| `failDetails[].id` | string | Tag ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

### Error Codes

| Code | Description |
|------|-------------|
| 43022 | Tag not found |

## tag-set

Set tags on materials. **Each operation replaces all existing tags** with the newly set tags.

| Flag | Description |
|------|-------------|
| `--materials-id <csv>` | Material IDs (required) |
| `--tags-id <csv>` | Tag IDs (will replace all existing tags) |

```bash
geelark-cli phone library tag-set --materials-id "mat1,mat2" --tags-id "tag1,tag2"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `failDetails[]` | array | Failure details (only present if some operations fail) |
| `failDetails[].id` | string | Material or Tag ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

### Error Codes

| Code | Description |
|------|-------------|
| 43022 | Tag not found |
| 60005 | Material not found |
