package workflow

import (
	"fmt"
	"os"

	"l/internal/domain"
)

// ListEntries reads the directory at dir and returns its entries.
// Hidden files (dot-prefixed) are excluded unless showHidden is true.
func (s *Service) ListEntries(dir string, showHidden bool) ([]domain.Entry, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory: %w", err)
	}

	entries := make([]domain.Entry, 0, len(dirEntries))
	for _, de := range dirEntries {
		if !showHidden && len(de.Name()) > 0 && de.Name()[0] == '.' {
			continue
		}
		info, err := de.Info()
		if err != nil {
			continue
		}
		t := domain.EntryTypeFile
		if de.IsDir() {
			t = domain.EntryTypeDir
		}
		entries = append(entries, domain.Entry{
			Name:     de.Name(),
			Type:     t,
			Size:     info.Size(),
			Modified: info.ModTime(),
		})
	}
	return entries, nil
}
