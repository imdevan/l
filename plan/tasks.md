# Context

This project converts the `list_all.sh` shell function into a Go CLI tool called `l`.
It is built on top of `go-cli-template`, which provides Cobra, Bubble Tea TUI, config management, and project scaffolding.

The tool lists directory contents in a styled table — a drop-in replacement for `ls` with sorting, filtering, color-coded output, and an optional interactive mode.

# Definitions

- **entry**: a file or directory item returned from reading a directory
- **sort key**: one of `name`, `type`, `size`, `modified`
- **age color**: color applied to the modified date based on how recent it is
- **interactive mode**: Bubble Tea inline TUI for filtering/navigating entries
- **config**: TOML config file managed by the existing `internal/config` package

---

# v0.1.0

## Feature 1: Directory listing core
- note: see config colors below for guidance on implementing color in the app implementation. 

colors should all be defined and routed through the config. use ansi colors by default but allow user to use hex as well. (same as curren implementation)

- [x] 1.1 Read current directory entries (name, type, size, modified)
- [x] 1.2 Render entries as a styled table using lipgloss
  - notes: columns — name, type, size, modified. Match the nushell table aesthetic from the plan example.
  - [x] 1.2.1 Color directory names (dir_color)
  - [x] 1.2.2 Color modified date by age (mod_newer → mod_new → mod_old → mod_older)
  - [x] 1.2.3 Color file size by bucket (file_sm → file_md → file_lg → file_xl)
  - [x] 1.2.4 Show "nothing here" message when directory is empty
- [x] 1.3 Show Nerd Font icons per entry type when `show_icons=true`
### config colors:
header_color=10
border_color=08
dir_color=12
file_color=07
# modified colors:
mod_newer_color = 10
mod_new_color = 02
mod_old_color = 03
mod_older_color = 04
modified_color = 08 # overrides other modified colors if present
# file size colorsj
file_sm=10
file_md=02
file_lg=03
file_xl=04
file_size = 08 # overrides other file colors if present


## Feature 2: Sorting

- [x] 2.1 Default sort by name (case-insensitive, natural order)
- [x] 2.2 `-m` flag: sort by modified (default descending)
- [x] 2.3 `-t` flag: sort by type (dirs first, then files)
- [x] 2.4 `-s` flag: sort by size (default descending)
- [x] 2.5 `-n` flag: sort by name (explicit; wins over other sort flags)
- [x] 2.6 `-r` flag: reverse sort order
  - notes: `-r` toggles the default direction, it does not force ascending. e.g. `-m` is desc by default; `-m -r` becomes ascending.
- [x] 2.7 `-f` only show files
- [x] 2.8 `-d` only show directories

## Feature 3: Filtering & navigation

- [x] 3.1 Optional positional arg as a filter query (substring match on name)
- [ ] 3.2 Directory hop: if query exactly matches a directory name, `cd` into it and list
  - [ ] 3.2.1 If query matches exactly one directory (fuzzy), hop into it
  - [ ] 3.2.2 Two positional args: first is the directory to hop, second is the filter query inside that directory
  - [ ] 3.2.3 Warn when second arg is ignored (ambiguous first arg, no hop occurred)
- [x] 3.3 `-a` flag: show hidden files (dotfiles)

## Feature 4: Config

- [x] 4.1 Wire new config fields into `internal/config`
  - notes: all fields listed in plan/simple.md under `# config`
- [x] 4.2 Add config fields to `example-config.toml` and `config_init.go`
- [x] 4.3 `default_flags` string: parsed and applied as if passed on the CLI
<!-- - [x] 4.4 `empty_dir_message` string: returned if directory is empty. default: "empty dir" -->

## Feature 5: Display options (config-driven)

- see config booleans below for default values 
- update default config. update docs

- [x] 5.1 `show_header` — toggle column header row
- [x] 5.2 `show_border` — toggle table border
- [x] 5.3 `show_type` — toggle type column
- [x] 5.4 `show_size` — toggle size column
- [x] 5.5 `show_modified` — toggle modified column
- [ ] 5.6 `show_bottom_header_for_large_returns` — repeat header at bottom when row count is large
- [ ] 5.7 `always_show_bottom_header` — always repeat header at bottom
- [x] 5.8 `show_permissions` — show permissions column when enabled
- [x] 5.9 `empty_dir_message` — message to show when directory is empty

### config booleans:
show_header=true
show_bottom_header_for_large_returns=true
always_show_bottom_header=false
show_border=true
show_permissions=false
show_modified=true
show_type=true
show_size=true
show_icons=true
- if true render with icons for file type
default_interactive=false
## Feature 6: Interactive mode

- [ ] 6.1 `/` key enters filter input; typing narrows the visible entries
- [ ] 6.2 `default_interactive=true` config option drops into interactive mode by default
- [ ] 6.3 Esc / Enter exits interactive mode
  - notes: use Bubble Tea inline mode (no full-screen takeover). See existing `internal/ui/list.go` for reference pattern.
