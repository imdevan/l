---
title: l
description: an ls replacement
---


<img width="1095" height="628" alt="screenshot-2026-06-11_21-59-29" src="https://github.com/user-attachments/assets/d3592213-ce7e-498b-9d95-3853dad6e13a" />


`l` is a fast, styled `ls` replacement with Nerd Font icons, color-coded output, and full TOML configuration.

Inspired by Nushell ls function. Re-implemented in Go, using cobra + lipgloss.

## Features

- Styled table output with color-coded names, sizes, and modified dates
- Nerd Font icons per file type (via [go-devicons](https://github.com/epilande/go-devicons))
- Sort by name, type, size, or modified date
- Filter by substring, show/hide hidden files, show only files or dirs
- Fully configurable via TOML — all colors, display options, and default flags

## Usage

```
l [filter] [flags]  # list current directory, optionally filtered by name

Sorting:
  -m                sort by modified (newest first)
  -s                sort by size (largest first)
  -t                sort by type (dirs first)
  -n                sort by name (wins over other sort flags)
  -r                reverse sort order

Filtering:
  -a                show hidden files (dotfiles)
  -f                only show files
  -d                only show directories

Config:
  -C                open config in editor
  --config-init     generate default config file
  -c <path>         use a specific config file

Other:
  -v                print version
  --completion zsh  print shell completion script (bash|zsh|fish|powershell)
  -h                help
```


## Configuration

```bash
l -c            # open or init config
```

```bash
l --config-init # init or reset config
```

See [contributing](https://devan.gg/l/contributing/) for more info.

## Installation

### Homebrew

```bash
brew install imdevan/l/l
```

### AUR

```bash
yay -S l
```

See [install](https://devan.gg/l/install/) for more info.

## Requirements

- [Nerd Font](https://www.nerdfonts.com/) — for icons (optional, disable with `show_icons = false`)

# Development

See [development](https://devan.gg/l/development/) for more information.
