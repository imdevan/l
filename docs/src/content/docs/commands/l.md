---
title: l
description: an ls replacement
---

an ls replacement

## Usage

```bash
l [alias]
l [command]
```

## Flags

| Flag | Type | Description |
|------|------|-------------|
| `-v, --version` | bool | print version information |
| `-c, --config` | bool | open config in editor |
| `-C, --config-init` | bool | generate default config file |
| `-y, --yes` | bool | skip overwrite confirmation |
| `-m, --sort-modified` | bool | sort by modified (newest first) |
| `-t, --sort-type` | bool | sort by type (dirs first) |
| `-s, --sort-size` | bool | sort by size (largest first) |
| `-n, --sort-name` | bool | sort by name (wins over other sort flags) |
| `-r, --reverse` | bool | reverse sort order |
| `-a, --all` | bool | show hidden files (dotfiles) |
| `-f, --files` | bool | only show files |
| `-d, --dirs` | bool | only show directories |

## Available Commands


## Source

See [root.go](https://github.com/imdevan/l//blob/main/cmd/l/root.go) for implementation details.
