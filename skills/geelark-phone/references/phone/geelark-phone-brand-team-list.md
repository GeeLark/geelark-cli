# phone brand-team-list

List team-uploaded cloud phone brands and models for a given Android version.

## Key Flags

| Flag | Description |
|------|-------------|
| `--android-ver <n>` | Android version (9/10/11/12/13/15) (required) |
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |

## Examples

```bash
geelark-cli phone brand-team-list --android-ver 10 --page 1 --page-size 10
```
