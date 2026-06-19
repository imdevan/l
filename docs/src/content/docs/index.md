---
title: l
description: an ls replacement
---


<img width="961" height="540" alt="l main" src="https://github.com/user-attachments/assets/82785068-de69-4740-b0b7-fa7210d33fd9" />


`l` is a fast, styled `ls` replacement with Nerd Font icons, color-coded output, and full TOML configuration.

Inspired by Nushell ls function. Re-implemented in Go, using cobra + lipgloss.

## Features

- Styled table output with color-coded names, sizes, and modified dates
- Nerd Font icons per file type (via [go-devicons](https://github.com/epilande/go-devicons))
- Sort by name, type, size, or modified date
- Filter by substring, show/hide hidden files, show only files or dirs
- Fully configurable via TOML — all colors, display options, and default flags

## Usage

```bash
$ l -h

Usage:
  l [filter] [flags]

Flags:
  -a, --all                   show hidden files (dotfiles)
  -c, --config                open config in editor
  -C, --config-init           generate default config file
  -d, --dirs                  only show directories
  -f, --files                 only show files
  -h, --help                  help for l
      --local-config string   config file path
  -r, --reverse               reverse sort order
  -m, --sort-modified         sort by modified (newest first)
  -n, --sort-name             sort by name (wins over other sort flags)
  -s, --sort-size             sort by size (largest first)
  -t, --sort-type             sort by type (dirs first)
  -v, --version               print version information
```

See [command](https://devan.gg/l/commands/l/) for more info.


## Configuration

```bash
l -c            # open or init config
```

```bash
l --config-init # init or reset config
```

See [configuration](https://devan.gg/l/configuration/) for more info.

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

See [contributing](https://devan.gg/l/contributing/) for more info.

# Additional Examples

### My current fav

<img width="960" height="540" alt="l fav" src="https://github.com/user-attachments/assets/2c9062ef-dcb6-484d-9642-2bcd9d3c8e98" />

### Mostly borderless with header

<img width="960" height="540" alt="l borderless" src="https://github.com/user-attachments/assets/772f3f97-fc7d-4be2-87d3-c1b36f307c62" />

### Borderless with permissions

<img width="960" height="540" alt="l minimal perms" src="https://github.com/user-attachments/assets/41438e75-a0e5-4111-991c-ce17d648d9f8" />

### Minimal

<img width="961" height="540" alt="l minimal" src="https://github.com/user-attachments/assets/2fbacd76-3764-457d-ae21-76b6c8321eb8" />






