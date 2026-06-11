package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"

	"l/internal/domain"
	"l/internal/utils"
)

// ManagerImpl loads and saves configuration files.
type ManagerImpl struct {
	cwd string
}

// NewManager returns a config manager rooted at the provided cwd.
func NewManager(cwd string) *ManagerImpl {
	return &ManagerImpl{cwd: cwd}
}

// LoadWithOverride loads config from a specific path, layered on defaults.
func (m *ManagerImpl) LoadWithOverride(path string) (domain.Config, error) {
	config := domain.DefaultConfig()
	if strings.TrimSpace(path) == "" {
		return m.Load()
	}
	partial, err := readConfig(path)
	if err != nil {
		return domain.Config{}, err
	}
	if partial != nil {
		applyPartial(&config, partial)
	}
	return config, nil
}

// Load reads config with precedence: defaults < global < local.
func (m *ManagerImpl) Load() (domain.Config, error) {
	config := domain.DefaultConfig()

	globalPath := utils.ConfigPathGlobal()
	if partial, err := readConfig(globalPath); err != nil {
		return domain.Config{}, err
	} else if partial != nil {
		applyPartial(&config, partial)
	}

	localPath := utils.ConfigPathLocal(m.cwd)
	if partial, err := readConfig(localPath); err != nil {
		return domain.Config{}, err
	} else if partial != nil {
		applyPartial(&config, partial)
	}

	return config, nil
}

// Save persists config to the global config path.
func (m *ManagerImpl) Save(config domain.Config) error {
	path := utils.ConfigPathGlobal()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Exists reports whether a local or global config file exists.
func (m *ManagerImpl) Exists() (bool, error) {
	globalPath := utils.ConfigPathGlobal()
	if exists, err := fileExists(globalPath); err != nil {
		return false, err
	} else if exists {
		return true, nil
	}
	localPath := utils.ConfigPathLocal(m.cwd)
	return fileExists(localPath)
}

type partialConfig struct {
	Editor               *string `toml:"editor"`
	Primary              *string `toml:"primary"`
	Secondary            *string `toml:"secondary"`
	Headings             *string `toml:"headings"`
	Text                 *string `toml:"text"`
	TextHighlight        *string `toml:"text_highlight"`
	DescriptionHighlight *string `toml:"description_highlight"`
	Tags                 *string `toml:"tags"`
	Flags                *string `toml:"flags"`
	Muted                *string `toml:"muted"`
	Accent               *string `toml:"accent"`
	Border               *string `toml:"border"`
	InteractiveDefault   *bool   `toml:"interactive_default"`
	ListSpacing          *string `toml:"list_spacing"`

	// Listing colors
	DirColor         *string `toml:"dir_color"`
	FileColor        *string `toml:"file_color"`
	TypeColor        *string `toml:"type_color"`
	PermissionsColor *string `toml:"permissions_color"`

	// Modified age colors
	ModNewerColor *string `toml:"mod_newer_color"`
	ModNewColor   *string `toml:"mod_new_color"`
	ModOldColor   *string `toml:"mod_old_color"`
	ModOlderColor *string `toml:"mod_older_color"`
	ModifiedColor *string `toml:"modified_color"`

	// File size colors
	FileSm   *string `toml:"file_sm"`
	FileMd   *string `toml:"file_md"`
	FileLg   *string `toml:"file_lg"`
	FileXl   *string `toml:"file_xl"`
	FileSize *string `toml:"file_size"`

	// Behavior
	DefaultFlags    *string `toml:"default_flags"`
	EmptyDirMessage *string `toml:"empty_dir_message"`

	// Display options
	ShowHeader                      *bool `toml:"show_header"`
	ShowBorder                      *bool `toml:"show_border"`
	ShowType                        *bool `toml:"show_type"`
	ShowSize                        *bool `toml:"show_size"`
	ShowModified                    *bool `toml:"show_modified"`
	ShowIcons                       *bool `toml:"show_icons"`
	ShowPermissions                 *bool `toml:"show_permissions"`
	ShowBottomHeaderForLargeReturns *bool `toml:"show_bottom_header_for_large_returns"`
	AlwaysShowBottomHeader          *bool `toml:"always_show_bottom_header"`
	DirectoriesFirst                *bool `toml:"directories_first"`
}

func readConfig(path string) (*partialConfig, error) {
	if exists, err := fileExists(path); err != nil || !exists {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var partial partialConfig
	if err := toml.Unmarshal(data, &partial); err != nil {
		return nil, err
	}
	return &partial, nil
}

func applyPartial(config *domain.Config, partial *partialConfig) {
	if partial.Editor != nil { config.Editor = *partial.Editor }
	if partial.Primary != nil { config.Primary = *partial.Primary }
	if partial.Secondary != nil { config.Secondary = *partial.Secondary }
	if partial.Headings != nil { config.Headings = *partial.Headings }
	if partial.Text != nil { config.Text = *partial.Text }
	if partial.TextHighlight != nil { config.TextHighlight = *partial.TextHighlight }
	if partial.DescriptionHighlight != nil { config.DescriptionHighlight = *partial.DescriptionHighlight }
	if partial.Tags != nil { config.Tags = *partial.Tags }
	if partial.Flags != nil { config.Flags = *partial.Flags }
	if partial.Muted != nil { config.Muted = *partial.Muted }
	if partial.Accent != nil { config.Accent = *partial.Accent }
	if partial.Border != nil { config.Border = *partial.Border }
	if partial.InteractiveDefault != nil { config.InteractiveDefault = *partial.InteractiveDefault }
	if partial.ListSpacing != nil { config.ListSpacing = *partial.ListSpacing }

	if partial.DirColor != nil { config.DirColor = *partial.DirColor }
	if partial.FileColor != nil { config.FileColor = *partial.FileColor }
	if partial.TypeColor != nil { config.TypeColor = *partial.TypeColor }
	if partial.PermissionsColor != nil { config.PermissionsColor = *partial.PermissionsColor }

	if partial.ModNewerColor != nil { config.ModNewerColor = *partial.ModNewerColor }
	if partial.ModNewColor != nil { config.ModNewColor = *partial.ModNewColor }
	if partial.ModOldColor != nil { config.ModOldColor = *partial.ModOldColor }
	if partial.ModOlderColor != nil { config.ModOlderColor = *partial.ModOlderColor }
	if partial.ModifiedColor != nil { config.ModifiedColor = *partial.ModifiedColor }

	if partial.FileSm != nil { config.FileSm = *partial.FileSm }
	if partial.FileMd != nil { config.FileMd = *partial.FileMd }
	if partial.FileLg != nil { config.FileLg = *partial.FileLg }
	if partial.FileXl != nil { config.FileXl = *partial.FileXl }
	if partial.FileSize != nil { config.FileSize = *partial.FileSize }

	if partial.DefaultFlags != nil { config.DefaultFlags = *partial.DefaultFlags }
	if partial.EmptyDirMessage != nil { config.EmptyDirMessage = *partial.EmptyDirMessage }

	if partial.ShowHeader != nil { config.ShowHeader = *partial.ShowHeader }
	if partial.ShowBorder != nil { config.ShowBorder = *partial.ShowBorder }
	if partial.ShowType != nil { config.ShowType = *partial.ShowType }
	if partial.ShowSize != nil { config.ShowSize = *partial.ShowSize }
	if partial.ShowModified != nil { config.ShowModified = *partial.ShowModified }
	if partial.ShowIcons != nil { config.ShowIcons = *partial.ShowIcons }
	if partial.ShowPermissions != nil { config.ShowPermissions = *partial.ShowPermissions }
	if partial.ShowBottomHeaderForLargeReturns != nil { config.ShowBottomHeaderForLargeReturns = *partial.ShowBottomHeaderForLargeReturns }
	if partial.AlwaysShowBottomHeader != nil { config.AlwaysShowBottomHeader = *partial.AlwaysShowBottomHeader }
	if partial.DirectoriesFirst != nil { config.DirectoriesFirst = *partial.DirectoriesFirst }
}

func expandPath(value string) string {
	expanded := os.ExpandEnv(value)
	if expanded == "" {
		return expanded
	}
	if expanded == "~" {
		if home, err := os.UserHomeDir(); err == nil {
			return home
		}
		return expanded
	}
	if strings.HasPrefix(expanded, "~/") {
		if home, err := os.UserHomeDir(); err == nil {
			return filepath.Join(home, strings.TrimPrefix(expanded, "~/"))
		}
	}
	return expanded
}

func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir(), nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
