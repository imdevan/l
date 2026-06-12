# l

`l` is a fast, styled `ls` replacement with Nerd Font icons, color-coded output, and full TOML configuration.

## Features

- Styled table output with color-coded names, sizes, and modified dates
- Nerd Font icons per file type (via [go-devicons](https://github.com/epilande/go-devicons))
- Sort by name, type, size, or modified date
- Filter by substring, show/hide hidden files, show only files or dirs
- Fully configurable via TOML — all colors, display options, and default flags

## Usage

```
l [filter] [flags] list current directory, optionally filtered by name

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

Configuration is stored at `$XDG_CONFIG_HOME/l/config.toml` (typically `~/.config/l/config.toml`).

Generate a config file with defaults:

```bash
l --config-init
```

Open config in your editor:

```bash
l -C
```

See [CONFIG.md](./CONFIG.md) for all options, or [example-config.toml](./example-config.toml) for a ready-to-copy example.

## Installation

See [INSTALL.md](./INSTALL.md).

## Requirements

- [Go](https://go.dev/) — build from source
- [Just](https://just.systems/) — run scripts
- Nerd Font — for icons (optional, disable with `show_icons = false`)

# Development
## Quick start

```bash
gh repo clone imdevan/l
cd l
just build-run
```

## Just scripts

```bash
just build           # Build binary
just build-run       # Build and run
just test            # Run tests
just install         # Install to /usr/local/bin
just watch           # Rebuild on file changes
```

