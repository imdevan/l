# context

this project is currently a go cli template that shows the current directory files.

convert /home/devy/dotfiles/functions/list_all.sh into go function

default output "l"

example:
```bash
Obsidian Vault/todo master  ?
❯ l

Default example based on nu shell 
╭──────────────────┬──────┬────────┬────────────────╮
│       name       │ type │  size  │    modified    │
├──────────────────┼──────┼────────┼────────────────┤
│ Fire Fix.md      │ file │  229 B │ 5 days ago     │
│ One day maybe    │ dir  │   22 B │ 4 weeks ago    │
│ plan             │ dir  │   44 B │ 12 minutes ago │
│ Project Ideas.md │ file │ 3.3 kB │ 12 minutes ago │
│ Recap            │ dir  │   26 B │ 4 weeks ago    │
│ Rewards.md       │ file │  134 B │ 2 days ago     │
│ The Wheel.md     │ file │  445 B │ 12 minutes ago │
│ This Week.md     │ file │  150 B │ 2 months ago   │
│ videos.md        │ file │ 1.3 kB │ 4 weeks ago    │
│ Want to Work.md  │ file │  117 B │ 2 weeks ago    │
╰──────────────────┴──────┴────────┴────────────────╯
```

# Sorting flags (precedence: -n wins if multiple are present)
-m sort="modified"
-t sort="type"
-s sort="size"
-n sort="name"

-r reverse
-a show_hidden

# config
config properties with their defaults

### config strings:
default_flags="" # default flags to use on all l calls
empty_dir_message="empty" # message to show when directory is empty

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


# interactive mode
 
- start typing to filter? or / to filter? 
  - / leaves room to use alpha for nav or hotkeys

questions:
should i include crud operations? 

