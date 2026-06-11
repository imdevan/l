package icon

// Icon represents a nerd font icon
type Icon string

const (
	// Bookmarks icon
	Bookmarks Icon = "\uf097"

	// Tmux icon
	Tmux Icon = "\uebc8"

	// Editor icons
	Nvim Icon = "\uf36f"
	Vim  Icon = "\uf36f"

	// File/Script icons
	File   Icon = "\uf15b" // Generic file icon
	Script Icon = "\ue691" // Script/terminal icon
	Shell  Icon = "\uf489" // Shell icon

	// Directory/entry type icons
	Dir     Icon = "\uf07b" // Folder icon
	FileMd  Icon = "\uf15b" // File icon (alias)
	Symlink Icon = "\uf481" // Symlink icon
)

// Get returns the icon as a string
func (i Icon) String() string {
	return string(i)
}

// GetEditorIcon returns the appropriate icon for a given editor
func GetEditorIcon(editor string) Icon {
	switch editor {
	case "nvim", "neovim":
		return Nvim
	case "vim":
		return Vim
	default:
		return ""
	}
}
