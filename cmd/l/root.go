package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"

	"l/internal/config"
	"l/internal/domain"
	pkg "l/internal/package"
	"l/internal/ui"
	"l/internal/workflow"
)

var (
	version = pkg.Version()
	name    = pkg.Name()
	short   = pkg.Short()
)

type rootOptions struct {
	localConfig string
	showVersion bool
	openConfig  bool
	initConfig  bool
	yes         bool
	completion  string
	sortByM     bool
	sortByT     bool
	sortByS     bool
	sortByN     bool
	reverse     bool
	showHidden  bool
	filesOnly   bool
	dirsOnly    bool
}

func (o *rootOptions) sortKey() workflow.SortKey {
	switch {
	case o.sortByN:
		return workflow.SortName
	case o.sortByM:
		return workflow.SortModified
	case o.sortByS:
		return workflow.SortSize
	case o.sortByT:
		return workflow.SortType
	default:
		return workflow.SortName
	}
}

func (o *rootOptions) anySortFlagSet() bool {
	return o.sortByM || o.sortByT || o.sortByS || o.sortByN || o.reverse
}

var rootCmd = newRootCmd()

func Execute() error {
	return rootCmd.Execute()
}

func newRootCmd() *cobra.Command {
	opts := &rootOptions{}
	cmd := &cobra.Command{
		Use:   name + " [filter]",
		Short: short,
		Long: `l — a styled directory listing tool

Usage:
  l [filter]          list current directory, optionally filtered by name
  l -m                sort by modified (newest first)
  l -s                sort by size (largest first)
  l -t                sort by type (dirs first)
  l -n                sort by name
  l -r                reverse sort order
  l -a                show hidden files

Config:
  l -c                open config in editor
  l -C                generate default config file
  l --local-config <path>  use specific config file

Other:
  l -v                print version
  l --completion zsh  print shell completion script`,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.showVersion {
				cmd.Printf("%s\n", resolvedVersion())
				return nil
			}
			if opts.completion != "" {
				return runCompletion(cmd, opts.completion)
			}
			if opts.initConfig {
				return runConfigInit(cmd, &configInitOptions{yes: opts.yes})
			}
			if opts.openConfig {
				return runConfig(cmd)
			}
			return runListing(cmd, opts, args)
		},
	}

	cmd.Flags().StringVar(&opts.localConfig, "local-config", "", "config file path")
	cmd.Flags().BoolVarP(&opts.showVersion, "version", "v", false, "print version information")
	cmd.Flags().BoolVarP(&opts.openConfig, "config", "c", false, "open config in editor")
	cmd.Flags().BoolVarP(&opts.initConfig, "config-init", "C", false, "generate default config file")
	cmd.Flags().BoolVarP(&opts.yes, "yes", "y", false, "skip overwrite confirmation")
	cmd.Flags().StringVar(&opts.completion, "completion", "", "print shell completion script (bash|zsh|fish|powershell)")
	cmd.Flags().BoolVarP(&opts.sortByM, "sort-modified", "m", false, "sort by modified (newest first)")
	cmd.Flags().BoolVarP(&opts.sortByT, "sort-type", "t", false, "sort by type (dirs first)")
	cmd.Flags().BoolVarP(&opts.sortByS, "sort-size", "s", false, "sort by size (largest first)")
	cmd.Flags().BoolVarP(&opts.sortByN, "sort-name", "n", false, "sort by name (wins over other sort flags)")
	cmd.Flags().BoolVarP(&opts.reverse, "reverse", "r", false, "reverse sort order")
	cmd.Flags().BoolVarP(&opts.showHidden, "all", "a", false, "show hidden files (dotfiles)")
	cmd.Flags().BoolVarP(&opts.filesOnly, "files", "f", false, "only show files")
	cmd.Flags().BoolVarP(&opts.dirsOnly, "dirs", "d", false, "only show directories")

	return cmd
}

func resolvedVersion() string {
	ver := version
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ver
	}
	if ver == "dev" && strings.TrimSpace(info.Main.Version) != "" && info.Main.Version != "(devel)" {
		ver = info.Main.Version
	}
	return ver
}

func runListing(cmd *cobra.Command, opts *rootOptions, args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	manager := config.NewManager(cwd)
	var cfg domain.Config
	if opts.localConfig != "" {
		cfg, err = manager.LoadWithOverride(opts.localConfig)
	} else {
		cfg, err = manager.Load()
	}
	if err != nil {
		cfg = domain.DefaultConfig()
	}

	if !opts.anySortFlagSet() && strings.TrimSpace(cfg.DefaultFlags) != "" {
		defaultArgs := strings.Fields(cfg.DefaultFlags)
		if err := cmd.Flags().Parse(defaultArgs); err != nil {
			return fmt.Errorf("invalid default_flags %q: %w", cfg.DefaultFlags, err)
		}
	}

	svc := workflow.New()
	entries, err := svc.ListEntries(cwd, opts.showHidden)
	if err != nil {
		return err
	}

	if opts.filesOnly {
		filtered := entries[:0]
		for _, e := range entries {
			if !e.IsDir() {
				filtered = append(filtered, e)
			}
		}
		entries = filtered
	} else if opts.dirsOnly {
		filtered := entries[:0]
		for _, e := range entries {
			if e.IsDir() {
				filtered = append(filtered, e)
			}
		}
		entries = filtered
	}

	if len(args) > 0 && strings.TrimSpace(args[0]) != "" {
		query := strings.ToLower(args[0])
		filtered := entries[:0]
		for _, e := range entries {
			if strings.Contains(strings.ToLower(e.Name), query) {
				filtered = append(filtered, e)
			}
		}
		entries = filtered
	}

	workflow.SortEntries(entries, opts.sortKey(), opts.reverse)

	if cfg.DirectoryPosition != "inline" && opts.sortKey() != workflow.SortType {
		workflow.SortDirs(entries, cfg.DirectoryPosition == "top")
	}

	theme := ui.ThemeFromConfig(cfg)
	fmt.Println(ui.RenderTable(entries, theme, ui.TableOptionsFromConfig(cfg)))
	return nil
}
