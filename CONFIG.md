# Configuration

Config file location: `~/.config/l/config.toml` 

(if you can't find  the config check `$XDG_CONFIG_HOME/l/config.toml`)

## Generate with defaults

```bash
l --config-init
```

## General

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `editor` | string | `nvim` | Editor opened by `l -C` |
| `interactive_default` | bool | `false` | Start in interactive mode by default |
| `default_flags` | string | — | Default CLI flags (e.g. `"-m -a"`) |
| `empty_dir_message` | string | `empty dir` | Message shown when directory is empty |

## Display

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `show_header` | bool | `true` | Show column header row |
| `show_border` | bool | `true` | Show table border |
| `show_type` | bool | `true` | Show file type column |
| `show_size` | bool | `true` | Show file size column |
| `show_modified` | bool | `true` | Show modified date column |
| `show_icons` | bool | `true` | Show Nerd Font icons |
| `show_permissions` | bool | `false` | Show permissions column |
| `show_bottom_header_for_large_returns` | bool | `true` | Repeat header at bottom when row count exceeds window height |
| `always_show_bottom_header` | bool | `false` | Always repeat header at bottom |
| `directory_position` | string | `inline` | Directory sort position: `top`, `bottom`, or `inline` (sorted with files) |

## Colors

Colors accept terminal color numbers, named colors, or hex values (e.g. `7`, `"red"`, `"#ff8800"`).

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `headings` | string | `15` | Column header color |
| `primary` | string | `02` | Primary accent color |
| `secondary` | string | `06` | Secondary accent color |
| `text` | string | `07` | Default text color |
| `text_highlight` | string | `06` | Highlighted text color |
| `description_highlight` | string | `05` | Highlighted description color |
| `muted` | string | `08` | Muted/dim text color |
| `border` | string | `08` | Table border color |
| `dir_color` | string | `12` | Directory name color |
| `file_color` | string | `07` | File name color |
| `type_color` | string | `08` | File type column color |
| `permissions_color` | string | `08` | Permissions column color |

### Modified age colors

Applied based on how recently a file was modified.

| Option | Default | When applied |
|--------|---------|--------------|
| `mod_newer_color` | `10` | < 1 hour ago |
| `mod_new_color` | `02` | < 1 day ago |
| `mod_old_color` | `03` | < 1 week ago |
| `mod_older_color` | `04` | < 4 weeks ago |
| `modified_color` | — | Overrides all of the above when set |

### File size colors

| Option | Default | When applied |
|--------|---------|--------------|
| `file_sm` | `10` | < 10 KB |
| `file_md` | `02` | < 100 KB |
| `file_lg` | `03` | < 1 MB |
| `file_xl` | `04` | ≥ 1 MB |
| `file_size` | — | Overrides all of the above when set |
