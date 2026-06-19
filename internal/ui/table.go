package ui

import (
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/epilande/go-devicons"
	"golang.org/x/term"

	"l/internal/domain"
)

// TableOptions controls which columns are rendered.
type TableOptions struct {
	ShowHeader                      bool
	ShowBorder                      bool
	ShowType                        bool
	ShowSize                        bool
	ShowModified                    bool
	ShowIcons                       bool
	ShowPermissions                 bool
	ShowBottomHeaderForLargeReturns bool
	AlwaysShowBottomHeader          bool
	EmptyMessage                    string
	MarginTop                       int
	MarginLeft                      int
}

// largeReturnThreshold is the row count at which the bottom header appears.
const largeReturnThreshold = 20

// getTerminalHeight returns the terminal height or -1 if not a terminal. Stubbable for tests.
var getTerminalHeight = func() int {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		_, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err == nil && h > 0 {
			return h
		}
	}
	if term.IsTerminal(int(os.Stderr.Fd())) {
		_, h, err := term.GetSize(int(os.Stderr.Fd()))
		if err == nil && h > 0 {
			return h
		}
	}
	if term.IsTerminal(int(os.Stdin.Fd())) {
		_, h, err := term.GetSize(int(os.Stdin.Fd()))
		if err == nil && h > 0 {
			return h
		}
	}
	return -1
}

// DefaultTableOptions returns sensible defaults matching the plan.
func DefaultTableOptions() TableOptions {
	return TableOptions{
		ShowHeader:                      true,
		ShowBorder:                      true,
		ShowType:                        true,
		ShowSize:                        true,
		ShowModified:                    true,
		ShowIcons:                       true,
		ShowPermissions:                 false,
		ShowBottomHeaderForLargeReturns: true,
		AlwaysShowBottomHeader:          false,
		EmptyMessage:                    "empty dir",
	}
}

// TableOptionsFromConfig builds TableOptions from a Config, with defaults.
func TableOptionsFromConfig(cfg domain.Config) TableOptions {
	opts := DefaultTableOptions()
	opts.ShowHeader = cfg.ShowHeader
	opts.ShowBorder = cfg.ShowBorder
	opts.ShowType = cfg.ShowType
	opts.ShowSize = cfg.ShowSize
	opts.ShowModified = cfg.ShowModified
	opts.ShowIcons = cfg.ShowIcons
	opts.ShowPermissions = cfg.ShowPermissions
	opts.ShowBottomHeaderForLargeReturns = cfg.ShowBottomHeaderForLargeReturns
	opts.AlwaysShowBottomHeader = cfg.AlwaysShowBottomHeader
	if strings.TrimSpace(cfg.EmptyDirMessage) != "" {
		opts.EmptyMessage = cfg.EmptyDirMessage
	}
	opts.MarginTop, opts.MarginLeft = domain.ParseMargin(cfg.Margin)
	return opts
}

// RenderTable renders entries as a styled lipgloss table.
func RenderTable(entries []domain.Entry, theme Theme, opts TableOptions) string {
	if len(entries) == 0 {
		out := lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(theme.Border).
			Foreground(theme.Muted).
			Render(opts.EmptyMessage)
		return applyMargin(out, opts)
	}

	headerStyle := lipgloss.NewStyle().Foreground(theme.Headings).Bold(true)
	b := lipgloss.NewStyle().Foreground(theme.Border) // border glyph style

	// Build rows
	rows := make([][]string, len(entries))
	for i, e := range entries {
		rows[i] = buildRow(e, theme, opts)
	}

	// Build header
	headers := buildHeaders(opts)

	// Compute column widths
	widths := columnWidths(headers, rows)

	var sb strings.Builder

	sep := b.Render("│")

	renderRow := func(cells []string, style func(string) string) string {
		var parts []string
		for i, cell := range cells {
			padded := padRight(cell, widths[i])
			parts = append(parts, style(padded))
		}
		if opts.ShowBorder {
			return sep + " " + strings.Join(parts, " "+sep+" ") + " " + sep
		}
		return "  " + strings.Join(parts, "   ") + "  "
	}

	borderLine := func(left, mid, right, horiz string) string {
		var parts []string
		for _, w := range widths {
			parts = append(parts, strings.Repeat(b.Render(horiz), w+2))
		}
		return b.Render(left) + strings.Join(parts, b.Render(mid)) + b.Render(right)
	}

	if opts.ShowBorder {
		sb.WriteString(borderLine("╭", "┬", "╮", "─") + "\n")
	}

	if opts.ShowHeader {
		sb.WriteString(renderRow(headers, func(s string) string { return headerStyle.Render(s) }) + "\n")
		sb.WriteString(borderLine("├", "┼", "┤", "─") + "\n")
	}

	for _, row := range rows {
		sb.WriteString(renderRow(row, func(s string) string { return s }) + "\n")
	}

	showBottomHeader := opts.AlwaysShowBottomHeader
	if !showBottomHeader && opts.ShowBottomHeaderForLargeReturns {
		extra := 0
		if opts.ShowHeader {
			extra += 2
		}
		if opts.ShowBorder {
			extra += 2
		}
		totalHeight := len(rows) + extra

		termHeight := getTerminalHeight()
		if termHeight > 0 {
			if totalHeight > termHeight {
				showBottomHeader = true
			}
		} else {
			if len(rows) >= largeReturnThreshold {
				showBottomHeader = true
			}
		}
	}

	if (opts.ShowHeader && showBottomHeader) || opts.AlwaysShowBottomHeader {
		sb.WriteString(borderLine("├", "┼", "┤", "─") + "\n")
		sb.WriteString(renderRow(headers, func(s string) string { return headerStyle.Render(s) }) + "\n")
	}

	if opts.ShowBorder {
		sb.WriteString(borderLine("╰", "┴", "╯", "─"))
	}

	return applyMargin(sb.String(), opts)
}

func applyMargin(s string, opts TableOptions) string {
	if opts.MarginTop == 0 && opts.MarginLeft == 0 {
		return s
	}
	return lipgloss.NewStyle().MarginTop(opts.MarginTop).MarginLeft(opts.MarginLeft).Render(s)
}

func buildHeaders(opts TableOptions) []string {
	var h []string
	if opts.ShowPermissions {
		h = append(h, "permissions")
	}
	h = append(h, "name")
	if opts.ShowType {
		h = append(h, "type")
	}
	if opts.ShowSize {
		h = append(h, "size")
	}
	if opts.ShowModified {
		h = append(h, "modified")
	}
	return h
}

func buildRow(e domain.Entry, theme Theme, opts TableOptions) []string {
	var row []string
	if opts.ShowPermissions {
		row = append(row, lipgloss.NewStyle().Foreground(theme.PermissionsColor).Render(e.Permissions))
	}
	row = append(row, coloredName(e, theme, opts.ShowIcons))
	if opts.ShowType {
		row = append(row, lipgloss.NewStyle().Foreground(theme.TypeColor).Render(string(e.Type)))
	}
	if opts.ShowSize {
		row = append(row, coloredSize(e, theme))
	}
	if opts.ShowModified {
		row = append(row, coloredModified(e.Modified, theme))
	}
	return row
}

func coloredName(e domain.Entry, theme Theme, showIcons bool) string {
	nameColor := theme.FileColor
	if e.IsDir() {
		nameColor = theme.DirColor
	}
	nameStyle := lipgloss.NewStyle().Foreground(nameColor)
	if e.IsDir() {
		nameStyle = nameStyle.Bold(true)
	}

	if showIcons && e.Info != nil {
		style := devicons.IconForInfo(e.Info)
		iconColor := lipgloss.Color(style.Color)
		iconGlyph := style.Icon
		if e.IsDir() {
			iconColor = nameColor
			iconGlyph = "\U000F0256" // nf-md-folder_outline
		}
		icon := lipgloss.NewStyle().Foreground(iconColor).Render(iconGlyph)
		return icon + " " + nameStyle.Render(e.Name)
	}
	return nameStyle.Render(e.Name)
}

func coloredSize(e domain.Entry, theme Theme) string {
	if e.IsDir() {
		return lipgloss.NewStyle().Foreground(theme.Muted).Render(humanize.Bytes(uint64(e.Size)))
	}
	size := e.Size
	var color lipgloss.Color
	switch {
	case size < 10_000:
		color = theme.FileSm
	case size < 100_000:
		color = theme.FileMd
	case size < 1_000_000:
		color = theme.FileLg
	default:
		color = theme.FileXl
	}
	return lipgloss.NewStyle().Foreground(color).Render(humanize.Bytes(uint64(size)))
}

func coloredModified(t time.Time, theme Theme) string {
	age := time.Since(t)
	var color lipgloss.Color
	switch {
	case age < time.Hour:
		color = theme.ModNewer
	case age < 24*time.Hour:
		color = theme.ModNew
	case age < 7*24*time.Hour:
		color = theme.ModOld
	case age < 28*24*time.Hour:
		color = theme.ModOlder
	default:
		color = theme.Muted
	}
	return lipgloss.NewStyle().Foreground(color).Render(humanize.Time(t))
}

// padRight pads a plain string (stripping ANSI for width calc) to width.
func padRight(s string, width int) string {
	visible := lipgloss.Width(s)
	if visible >= width {
		return s
	}
	return s + strings.Repeat(" ", width-visible)
}

func columnWidths(headers []string, rows [][]string) []int {
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = lipgloss.Width(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				w := lipgloss.Width(cell)
				if w > widths[i] {
					widths[i] = w
				}
			}
		}
	}
	return widths
}
