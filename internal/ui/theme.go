package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"l/internal/domain"
)

// Theme holds configurable colors for UI output.
type Theme struct {
	Headings             lipgloss.Color
	Primary              lipgloss.Color
	Secondary            lipgloss.Color
	Text                 lipgloss.Color
	TextHighlight        lipgloss.Color
	DescriptionHighlight lipgloss.Color
	Tags                 lipgloss.Color
	Flags                lipgloss.Color
	Muted                lipgloss.Color
	Border               lipgloss.Color

	// Listing colors
	DirColor         lipgloss.Color
	FileColor        lipgloss.Color
	TypeDirColor     lipgloss.Color
	TypeFileColor    lipgloss.Color
	PermissionsColor lipgloss.Color
	PermReadColor    lipgloss.Color
	PermWriteColor   lipgloss.Color
	PermExecColor    lipgloss.Color
	PermNoneColor    lipgloss.Color
	PermDirColor     lipgloss.Color

	// Modified age colors
	ModNewer  lipgloss.Color
	ModNew    lipgloss.Color
	ModOld    lipgloss.Color
	ModOlder  lipgloss.Color
	ModOldest lipgloss.Color

	// File size colors
	FileSm lipgloss.Color
	FileMd lipgloss.Color
	FileLg lipgloss.Color
	FileXl lipgloss.Color
}

// ThemeFromConfig builds a theme with safe fallbacks.
func ThemeFromConfig(cfg domain.Config) Theme {
	// mod colors: modified_color overrides individual bands when set
	modOverride := strings.TrimSpace(cfg.ModifiedColor)
	modNewer  := resolveColor(resolveFallback(modOverride, cfg.ModNewerColor), "10")
	modNew    := resolveColor(resolveFallback(modOverride, cfg.ModNewColor), "02")
	modOld    := resolveColor(resolveFallback(modOverride, cfg.ModOldColor), "03")
	modOlder  := resolveColor(resolveFallback(modOverride, cfg.ModOlderColor), "09")
	modOldest := resolveColor(resolveFallback(modOverride, cfg.ModOldestColor), "01")

	// size colors: file_size overrides individual bands when set
	sizeOverride := strings.TrimSpace(cfg.FileSize)
	fileSm := resolveColor(resolveFallback(sizeOverride, cfg.FileSm), "10")
	fileMd := resolveColor(resolveFallback(sizeOverride, cfg.FileMd), "02")
	fileLg := resolveColor(resolveFallback(sizeOverride, cfg.FileLg), "03")
	fileXl := resolveColor(resolveFallback(sizeOverride, cfg.FileXl), "01")

	return Theme{
		Headings:             resolveColor(cfg.Headings, "15"),
		Primary:              resolveColor(cfg.Primary, "02"),
		Secondary:            resolveColor(cfg.Secondary, "06"),
		Text:                 resolveColor(cfg.Text, "07"),
		TextHighlight:        resolveColor(resolveFallback(cfg.TextHighlight, cfg.Secondary), "06"),
		DescriptionHighlight: resolveColor(resolveFallback(cfg.DescriptionHighlight, cfg.Secondary), "06"),
		Tags:                 resolveColor(resolveFallback(cfg.Tags, cfg.Accent), "13"),
		Flags:                resolveColor(cfg.Flags, "12"),
		Muted:                resolveColor(cfg.Muted, "08"),
		Border:               resolveColor(cfg.Border, "08"),

		DirColor:         resolveColor(resolveFallback(cfg.DirColor, cfg.Primary), "12"),
		FileColor:        resolveColor(resolveFallback(cfg.FileColor, cfg.Text), "07"),
		TypeDirColor:     resolveColor(resolveFallback(cfg.TypeColor, cfg.DirColor, cfg.Primary), "12"),
		TypeFileColor:    resolveColor(resolveFallback(cfg.TypeColor, cfg.FileColor, cfg.Text), "07"),
		PermissionsColor: resolveColor(resolveFallback(cfg.PermissionsColor, cfg.Muted), "08"),
		PermReadColor:    resolveColor(resolveFallback(cfg.PermissionsColor, cfg.PermReadColor), "03"),
		PermWriteColor:   resolveColor(resolveFallback(cfg.PermissionsColor, cfg.PermWriteColor), "01"),
		PermExecColor:    resolveColor(resolveFallback(cfg.PermissionsColor, cfg.PermExecColor), "02"),
		PermNoneColor:    resolveColor(resolveFallback(cfg.PermissionsColor, cfg.PermNoneColor), "08"),
		PermDirColor:     resolveColor(resolveFallback(cfg.PermissionsColor, cfg.PermDirColor), "12"),

		ModNewer:  modNewer,
		ModNew:    modNew,
		ModOld:    modOld,
		ModOlder:  modOlder,
		ModOldest: modOldest,

		FileSm: fileSm,
		FileMd: fileMd,
		FileLg: fileLg,
		FileXl: fileXl,
	}
}

func resolveColor(value, fallback string) lipgloss.Color {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		trimmed = fallback
	}
	return lipgloss.Color(trimmed)
}

func resolveFallback(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}
