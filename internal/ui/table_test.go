package ui

import (
	"strings"
	"testing"
	"time"

	"l/internal/domain"
)

func TestRenderTable_BottomHeaderForLargeReturns(t *testing.T) {
	// Restore original getTerminalHeight
	origGetTerminalHeight := getTerminalHeight
	defer func() { getTerminalHeight = origGetTerminalHeight }()

	entries := []domain.Entry{
		{Name: "file1.txt", Size: 100, Modified: time.Now(), Type: domain.EntryTypeFile},
		{Name: "file2.txt", Size: 200, Modified: time.Now(), Type: domain.EntryTypeFile},
		{Name: "file3.txt", Size: 300, Modified: time.Now(), Type: domain.EntryTypeFile},
	}

	theme := Theme{
		Headings: "15",
		Border:   "08",
	}

	t.Run("never repeat when show_header is false", func(t *testing.T) {
		getTerminalHeight = func() int { return 2 } // tiny terminal
		opts := DefaultTableOptions()
		opts.ShowHeader = false
		opts.ShowBottomHeaderForLargeReturns = true
		opts.AlwaysShowBottomHeader = false

		result := RenderTable(entries, theme, opts)
		// Header shouldn't be rendered at all
		if strings.Contains(result, "name") {
			t.Errorf("expected header not to be printed, got: %s", result)
		}
	})

	t.Run("repeat when total height exceeds terminal height", func(t *testing.T) {
		// totalHeight of 3 entries with header and border is: 3 + 2 (header) + 2 (border) = 7
		getTerminalHeight = func() int { return 6 } // terminal is smaller than 7
		opts := DefaultTableOptions()
		opts.ShowHeader = true
		opts.ShowBorder = true
		opts.ShowBottomHeaderForLargeReturns = true
		opts.AlwaysShowBottomHeader = false

		result := RenderTable(entries, theme, opts)
		// Should repeat header at bottom, meaning "name" appears twice.
		if count := strings.Count(result, "name"); count != 2 {
			t.Errorf("expected header 'name' to appear 2 times, got %d. Result:\n%s", count, result)
		}
	})

	t.Run("do not repeat when total height is within terminal height", func(t *testing.T) {
		getTerminalHeight = func() int { return 10 } // terminal is larger than 7
		opts := DefaultTableOptions()
		opts.ShowHeader = true
		opts.ShowBorder = true
		opts.ShowBottomHeaderForLargeReturns = true
		opts.AlwaysShowBottomHeader = false

		result := RenderTable(entries, theme, opts)
		// Should only show header at top, meaning "name" appears once.
		if count := strings.Count(result, "name"); count != 1 {
			t.Errorf("expected header 'name' to appear 1 time, got %d. Result:\n%s", count, result)
		}
	})

	t.Run("fallback to largeReturnThreshold when not a terminal", func(t *testing.T) {
		getTerminalHeight = func() int { return -1 } // not a terminal
		opts := DefaultTableOptions()
		opts.ShowHeader = true
		opts.ShowBottomHeaderForLargeReturns = true
		opts.AlwaysShowBottomHeader = false

		// Fewer than largeReturnThreshold (20)
		result := RenderTable(entries, theme, opts)
		if count := strings.Count(result, "name"); count != 1 {
			t.Errorf("expected header 'name' to appear 1 time for fallback threshold with few entries, got %d", count)
		}

		// Make 20 entries
		largeEntries := make([]domain.Entry, 20)
		for i := 0; i < 20; i++ {
			largeEntries[i] = domain.Entry{Name: "file.txt", Size: 100, Modified: time.Now(), Type: domain.EntryTypeFile}
		}
		resultLarge := RenderTable(largeEntries, theme, opts)
		if count := strings.Count(resultLarge, "name"); count != 2 {
			t.Errorf("expected header 'name' to appear 2 times for fallback threshold with 20 entries, got %d", count)
		}
	})

	t.Run("always repeat when always_show_bottom_header is true", func(t *testing.T) {
		getTerminalHeight = func() int { return 100 } // very large terminal
		opts := DefaultTableOptions()
		opts.ShowHeader = true
		opts.AlwaysShowBottomHeader = true

		result := RenderTable(entries, theme, opts)
		if count := strings.Count(result, "name"); count != 2 {
			t.Errorf("expected header 'name' to appear 2 times when always_show_bottom_header is true, got %d. Result:\n%s", count, result)
		}
	})

	t.Run("repeat at bottom when show_header is false but always_show_bottom_header is true", func(t *testing.T) {
		opts := DefaultTableOptions()
		opts.ShowHeader = false
		opts.AlwaysShowBottomHeader = true

		result := RenderTable(entries, theme, opts)
		// Header should appear exactly once (at the bottom)
		if count := strings.Count(result, "name"); count != 1 {
			t.Errorf("expected header 'name' to appear exactly 1 time (at bottom), got %d. Result:\n%s", count, result)
		}
	})
}
