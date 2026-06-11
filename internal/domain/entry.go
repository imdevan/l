package domain

import (
	"os"
	"time"
)

// EntryType indicates whether an entry is a file or directory.
type EntryType string

const (
	EntryTypeDir  EntryType = "dir"
	EntryTypeFile EntryType = "file"
)

// Entry represents a single directory listing item.
type Entry struct {
	Name        string
	Type        EntryType
	Size        int64
	Modified    time.Time
	Permissions string    // e.g. "drwxr-xr-x"
	Info        os.FileInfo // used for icon resolution
}

func (e Entry) IsDir() bool { return e.Type == EntryTypeDir }
