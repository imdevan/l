package workflow

import (
	"sort"
	"strings"

	"l/internal/domain"
)

// SortKey identifies the field to sort by.
type SortKey string

const (
	SortName     SortKey = "name"
	SortModified SortKey = "modified"
	SortSize     SortKey = "size"
	SortType     SortKey = "type"
)

// SortEntries sorts entries in place.
// Modified and size default to descending; name and type default to ascending.
// The reverse flag toggles that default direction (XOR).
func SortEntries(entries []domain.Entry, key SortKey, reverse bool) {
	defaultDesc := key == SortModified || key == SortSize
	descending := defaultDesc != reverse // XOR

	sort.SliceStable(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		var less bool
		switch key {
		case SortModified:
			less = a.Modified.Before(b.Modified)
		case SortSize:
			less = a.Size < b.Size
		case SortType:
			// dirs first, then by name
			if a.IsDir() != b.IsDir() {
				less = a.IsDir()
			} else {
				less = naturalLess(a.Name, b.Name)
			}
		default: // SortName
			less = naturalLess(a.Name, b.Name)
		}
		if descending {
			return !less
		}
		return less
	})
}

// naturalLess does case-insensitive comparison.
func naturalLess(a, b string) bool {
	return strings.ToLower(a) < strings.ToLower(b)
}
