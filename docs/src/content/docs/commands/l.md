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
| `-c, --config` | string | config file path |
| `-v, --version` | bool | print version information |
| `-C, --Config` | bool | open config in editor |
| `-m, --sort-modified` | bool | sort by modified (newest first) |
| `-t, --sort-type` | bool | sort by type (dirs first) |
| `-s, --sort-size` | bool | sort by size (largest first) |
| `-n, --sort-name` | bool | sort by name (wins over other sort flags) |
| `-r, --reverse` | bool | reverse sort order |
| `-a, --all` | bool | show hidden files (dotfiles) |
| `-f, --files` | bool | only show files |
| `-d, --dirs` | bool | only show directories |

## Available Commands

- [`completion`](/commands/completion) - 
- [`config`](/commands/config) - View or edit configuration
- [`config init`](/commands/config-init) - Generate a default config file

## Source

See [root.go](https://github.com/imdevan/l//blob/main/cmd/l/root.go) for implementation details.
