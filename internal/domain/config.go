package domain

import (
	"os"
	"path/filepath"
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
	ModNewerColor string `toml:"mod_newer_color"`
	ModNewColor   string `toml:"mod_new_color"`
	ModOldColor   string `toml:"mod_old_color"`
	ModOlderColor string `toml:"mod_older_color"`
	ModifiedColor string `toml:"modified_color"` // overrides all mod colors when set

	// File size colors
	FileSm   string `toml:"file_sm"`
	FileMd   string `toml:"file_md"`
	FileLg   string `toml:"file_lg"`
	FileXl   string `toml:"file_xl"`
	FileSize string `toml:"file_size"` // overrides all size colors when set

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
	DirectoryPosition             string `toml:"directory_position"`

	// Permissions color
	PermissionsColor string `toml:"permissions_color"`
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
		TypeColor: "08",

		ModNewerColor: "10",
		ModNewColor:   "02",
		ModOldColor:   "03",
		ModOlderColor: "04",

		FileSm: "10",
		FileMd: "02",
		FileLg: "03",
		FileXl: "04",

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

		PermissionsColor: "08",
	}
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
