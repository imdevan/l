package domain

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Config describes the resolved configuration.
type Config struct {
	Editor               string `toml:"editor"`
	Primary              string `toml:"primary"`
	Secondary            string `toml:"secondary"`
	Headings             string `toml:"headings"`
	Text                 string `toml:"text"`
	TextHighlight        string `toml:"text_highlight"`
	DescriptionHighlight string `toml:"description_highlight"`
	Tags                 string `toml:"tags"`
	Flags                string `toml:"flags"`
	Muted                string `toml:"muted"`
	Accent               string `toml:"accent"`
	Border               string `toml:"border"`
	InteractiveDefault   bool   `toml:"interactive_default"`
	ListSpacing          string `toml:"list_spacing"`

	// Listing colors
	DirColor  string `toml:"dir_color"`
	FileColor string `toml:"file_color"`
	TypeColor string `toml:"type_color"`

	// Modified age colors
	ModNewerColor  string `toml:"mod_newer_color"`
	ModNewColor    string `toml:"mod_new_color"`
	ModOldColor    string `toml:"mod_old_color"`
	ModOlderColor  string `toml:"mod_older_color"`
	ModOldestColor string `toml:"mod_oldest_color"`
	ModifiedColor  string `toml:"modified_color"` // overrides all mod colors when set

	// File size colors
	FileSm   string `toml:"file_sm"`
	FileMd   string `toml:"file_md"`
	FileLg   string `toml:"file_lg"`
	FileXl   string `toml:"file_xl"`
	FileSize string `toml:"file_size"` // overrides all size colors when set

	// Margin: "#" sets top margin; "#,#" sets top and left margin
	Margin string `toml:"margin"`

	// Default CLI flags applied before user-supplied flags
	DefaultFlags string `toml:"default_flags"`

	// Message shown when a directory is empty
	EmptyDirMessage string `toml:"empty_dir_message"`

	// Display options
	ShowHeader                    bool `toml:"show_header"`
	ShowBorder                    bool `toml:"show_border"`
	ShowType                      bool `toml:"show_type"`
	ShowSize                      bool `toml:"show_size"`
	ShowModified                  bool `toml:"show_modified"`
	ShowIcons                     bool `toml:"show_icons"`
	ShowPermissions               bool `toml:"show_permissions"`
	ShowBottomHeaderForLargeReturns bool `toml:"show_bottom_header_for_large_returns"`
	AlwaysShowBottomHeader        bool `toml:"always_show_bottom_header"`
	ShowDirSize                   bool `toml:"show_dir_size"`
	DirectoryPosition             string `toml:"directory_position"`

	// Permissions color
	PermissionsColor string `toml:"permissions_color"` // overrides all perm colors when set
	PermReadColor    string `toml:"perm_read_color"`
	PermWriteColor   string `toml:"perm_write_color"`
	PermExecColor    string `toml:"perm_exec_color"`
	PermNoneColor    string `toml:"perm_none_color"`
	PermDirColor     string `toml:"perm_dir_color"`
}

// DefaultConfig returns the default configuration values.
func DefaultConfig() Config {
	return Config{
		Editor:               "nvim",
		Headings:             "15",
		Primary:              "02",
		Secondary:            "06",
		Text:                 "07",
		TextHighlight:        "06",
		DescriptionHighlight: "05",
		Tags:                 "13",
		Flags:                "12",
		Muted:                "08",
		Accent:               "13",
		Border:               "08",
		InteractiveDefault:   true,
		ListSpacing:          "space",

		DirColor:  "12",
		FileColor: "07",
		TypeColor: "",

		ModNewerColor:  "10",
		ModNewColor:    "02",
		ModOldColor:    "03",
		ModOlderColor:  "09",
		ModOldestColor: "01",

		FileSm: "10",
		FileMd: "02",
		FileLg: "03",
		FileXl: "01",

		EmptyDirMessage: "empty dir",

		ShowHeader:                      true,
		ShowBorder:                      true,
		ShowType:                        true,
		ShowSize:                        true,
		ShowModified:                    true,
		ShowIcons:                       true,
		ShowPermissions:                 false,
		ShowBottomHeaderForLargeReturns: true,
		AlwaysShowBottomHeader:          false,
		DirectoryPosition:               "inline",

		PermissionsColor: "",
		PermReadColor:    "03",
		PermWriteColor:   "01",
		PermExecColor:    "02",
		PermNoneColor:    "08",
		PermDirColor:     "12",
	}
}

// ParseMargin parses a margin string into (top, left) values.
// Format: "#" → left=# top=0; "#,#" → left=first top=second.
func ParseMargin(s string) (top, left int) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, 0
	}
	parts := strings.SplitN(s, ",", 2)
	if v, err := strconv.Atoi(strings.TrimSpace(parts[0])); err == nil && v >= 0 {
		left = v
	}
	if len(parts) == 2 {
		if v, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil && v >= 0 {
			top = v
		}
	}
	return top, left
}

func xdgHome(envKey, fallbackSuffix string) string {
	if value := os.Getenv(envKey); value != "" {
		return value
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, fallbackSuffix)
}
