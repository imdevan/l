package ui

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"

	"l/internal/adapters/icon"
	"l/internal/domain"
)

// TableOptions controls which columns are rendered.
type TableOptions struct {
	ShowHeader   bool
	ShowBorder   bool
	ShowType     bool
	ShowSize     bool
	ShowModified bool
	ShowIcons    bool
}

// DefaultTableOptions returns sensible defaults matching the plan.
func DefaultTableOptions() TableOptions {
	return TableOptions{
		ShowHeader:   true,
		ShowBorder:   true,
		ShowType:     true,
		ShowSize:     true,
		ShowModified: true,
		ShowIcons:    true,
	}
}

// RenderTable renders entries as a styled lipgloss table.
func RenderTable(entries []domain.Entry, theme Theme, opts TableOptions) string {
	if len(entries) == 0 {
		return lipgloss.NewStyle().Foreground(theme.Muted).Render("🧙🏼‍♂️ nothing here")
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

	if opts.ShowBorder {
		sb.WriteString(borderLine("╰", "┴", "╯", "─"))
	}

	return sb.String()
}

func buildHeaders(opts TableOptions) []string {
	h := []string{"name"}
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
	name := coloredName(e, theme, opts.ShowIcons)
	row := []string{name}
	if opts.ShowType {
		row = append(row, lipgloss.NewStyle().Foreground(theme.Muted).Render(string(e.Type)))
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
	var prefix string
	if showIcons {
		if e.IsDir() {
			prefix = icon.Dir.String() + " "
		} else {
			prefix = icon.File.String() + " "
		}
	}
	if e.IsDir() {
		return lipgloss.NewStyle().Foreground(theme.DirColor).Bold(true).Render(prefix + e.Name)
	}
	return lipgloss.NewStyle().Foreground(theme.FileColor).Render(prefix + e.Name)
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
