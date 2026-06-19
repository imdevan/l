---
title: l
description: an ls replacement
---

`l` is an incredibly simple tool with only one command: `l`.

`l` lists the contents of a directory in styled table.

## Usage

```bash
l [filter] [flags]
```

## Flags

### Sorting

| Flag | Type | Description |
|------|------|-------------|
| `-m, --sort-modified` | bool | sort by modified (newest first) |
| `-t, --sort-type` | bool | sort by type (dirs first) |
| `-s, --sort-size` | bool | sort by size (largest first) |
| `-n, --sort-name` | bool | sort by name (wins over other sort flags) |
| `-r, --reverse` | bool | reverse sort order |

### Filtering

| Flag | Type | Description |
|------|------|-------------|
| `-a, --all` | bool | show hidden files (dotfiles) |
| `-f, --files` | bool | only show files |
| `-d, --dirs` | bool | only show directories |
| `-p, --permissions` | bool | show permissions column |

### Config

| Flag | Type | Description |
|------|------|-------------|
| `-c, --config` | bool | open config in editor |
| `-C, --config-init` | bool | generate default config file |
| `--local-config` | string | config file path |

### Metadata

| Flag | Type | Description |
|------|------|-------------|
| `-v, --version` | bool | print version information |


## Source

See [root.go](https://github.com/imdevan/l//blob/main/cmd/l/root.go) for implementation details.
