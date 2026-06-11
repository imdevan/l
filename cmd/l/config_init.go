package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"l/internal/adapters/editor"
	"l/internal/config"
	"l/internal/domain"
	"l/internal/utils"
)

type configInitOptions struct {
	force        bool
	openInEditor bool
}

func newConfigInitCmd() *cobra.Command {
	opts := &configInitOptions{}
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Generate a default config file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConfigInit(cmd, opts)
		},
	}
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "overwrite existing config")
	cmd.Flags().BoolVarP(&opts.openInEditor, "editor", "e", false, "open config in editor after creation")
	return cmd
}

func runConfigInit(cmd *cobra.Command, opts *configInitOptions) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	manager := config.NewManager(cwd)
	exists, err := manager.Exists()
	if err != nil {
		return err
	}
	if exists && !opts.force {
		return fmt.Errorf("config already exists at %s (use --force to overwrite)", utils.ConfigPathGlobal())
	}
	cfg := domain.DefaultConfig()
	path := utils.ConfigPathGlobal()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	content := renderConfigTemplate(cfg)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return err
	}
	if opts.openInEditor {
		editorAdapter := editor.New(cfg.Editor)
		if err := editorAdapter.Open(path); err != nil {
			return err
		}
	}
	cmd.Printf("Wrote config to %s\n", utils.ConfigPathGlobal())
	return nil
}

func renderConfigTemplate(cfg domain.Config) string {
	var b strings.Builder
	b.WriteString("# General\n")
	b.WriteString(fmt.Sprintf("editor = %q\n", cfg.Editor))
	b.WriteString("\n# CLI behavior\n")
	b.WriteString(fmt.Sprintf("interactive_default = %t\n", cfg.InteractiveDefault))
	b.WriteString("# default_flags = \"\"\n")
	b.WriteString("# empty_dir_message = \"empty dir\"\n")
	b.WriteString("\n# Colors\n")
	b.WriteString("# Colors support named, numeric, or hex values (ex: 7, 13, \"#ff8800\").\n")
	b.WriteString(fmt.Sprintf("headings = %q\n", cfg.Headings))
	b.WriteString(fmt.Sprintf("primary = %q\n", cfg.Primary))
	b.WriteString(fmt.Sprintf("secondary = %q\n", cfg.Secondary))
	b.WriteString(fmt.Sprintf("text = %q\n", cfg.Text))
	b.WriteString(fmt.Sprintf("text_highlight = %q\n", cfg.TextHighlight))
	b.WriteString(fmt.Sprintf("description_highlight = %q\n", cfg.DescriptionHighlight))
	b.WriteString(fmt.Sprintf("tags = %q\n", cfg.Tags))
	b.WriteString(fmt.Sprintf("flags = %q\n", cfg.Flags))
	b.WriteString(fmt.Sprintf("muted = %q\n", cfg.Muted))
	b.WriteString(fmt.Sprintf("border = %q\n", cfg.Border))
	b.WriteString("\n# Listing colors\n")
	b.WriteString(fmt.Sprintf("dir_color = %q\n", cfg.DirColor))
	b.WriteString(fmt.Sprintf("file_color = %q\n", cfg.FileColor))
	b.WriteString(fmt.Sprintf("type_color = %q\n", cfg.TypeColor))
	b.WriteString("\n# Modified age colors (mod_newer < 1h, mod_new < 1d, mod_old < 1w, mod_older < 4w)\n")
	b.WriteString(fmt.Sprintf("mod_newer_color = %q\n", cfg.ModNewerColor))
	b.WriteString(fmt.Sprintf("mod_new_color = %q\n", cfg.ModNewColor))
	b.WriteString(fmt.Sprintf("mod_old_color = %q\n", cfg.ModOldColor))
	b.WriteString(fmt.Sprintf("mod_older_color = %q\n", cfg.ModOlderColor))
	b.WriteString("# modified_color = \"\"  # overrides all mod colors when set\n")
	b.WriteString("\n# File size colors (file_sm < 10KB, file_md < 100KB, file_lg < 1MB, file_xl >= 1MB)\n")
	b.WriteString(fmt.Sprintf("file_sm = %q\n", cfg.FileSm))
	b.WriteString(fmt.Sprintf("file_md = %q\n", cfg.FileMd))
	b.WriteString(fmt.Sprintf("file_lg = %q\n", cfg.FileLg))
	b.WriteString(fmt.Sprintf("file_xl = %q\n", cfg.FileXl))
	b.WriteString("# file_size = \"\"  # overrides all size colors when set\n")
	return b.String()
}
