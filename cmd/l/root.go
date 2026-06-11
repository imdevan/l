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
	configPath  string
	showVersion bool
	sortByM     bool // -m modified
	sortByT     bool // -t type
	sortByS     bool // -s size
	sortByN     bool // -n name (wins)
	reverse     bool // -r
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

// anySortFlagSet reports whether the user explicitly passed a sort/reverse flag.
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
		Use:   name,
		Short: short,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.showVersion {
				cmd.Printf("%s\n", resolvedVersion())
				return nil
			}
			return runListing(cmd, opts, args)
		},
	}

	cmd.Flags().StringVarP(&opts.configPath, "config", "c", "", "config file path")
	cmd.Flags().BoolVarP(&opts.showVersion, "version", "v", false, "print version information")
	cmd.Flags().BoolVarP(&opts.sortByM, "sort-modified", "m", false, "sort by modified (newest first)")
	cmd.Flags().BoolVarP(&opts.sortByT, "sort-type", "t", false, "sort by type (dirs first)")
	cmd.Flags().BoolVarP(&opts.sortByS, "sort-size", "s", false, "sort by size (largest first)")
	cmd.Flags().BoolVarP(&opts.sortByN, "sort-name", "n", false, "sort by name (wins over other sort flags)")
	cmd.Flags().BoolVarP(&opts.reverse, "reverse", "r", false, "reverse sort order")

	cmd.AddCommand(newConfigCmd())
	cmd.AddCommand(newCompletionCmd())

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
	if opts.configPath != "" {
		cfg, err = manager.LoadWithOverride(opts.configPath)
	} else {
		cfg, err = manager.Load()
	}
	if err != nil {
		cfg = domain.DefaultConfig()
	}

	// Apply default_flags from config when no sort/behavior flags were explicitly passed.
	if !opts.anySortFlagSet() && strings.TrimSpace(cfg.DefaultFlags) != "" {
		defaultArgs := strings.Fields(cfg.DefaultFlags)
		if err := cmd.Flags().Parse(defaultArgs); err != nil {
			return fmt.Errorf("invalid default_flags %q: %w", cfg.DefaultFlags, err)
		}
	}

	svc := workflow.New()
	entries, err := svc.ListEntries(cwd, false)
	if err != nil {
		return err
	}

	workflow.SortEntries(entries, opts.sortKey(), opts.reverse)

	theme := ui.ThemeFromConfig(cfg)
	fmt.Println(ui.RenderTable(entries, theme, ui.TableOptionsFromConfig(cfg)))
	return nil
}
