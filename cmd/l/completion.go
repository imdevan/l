package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func runCompletion(cmd *cobra.Command, shell string) error {
	switch strings.ToLower(strings.TrimSpace(shell)) {
	case "bash":
		return cmd.Root().GenBashCompletion(cmd.OutOrStdout())
	case "zsh":
		return cmd.Root().GenZshCompletion(cmd.OutOrStdout())
	case "fish":
		return cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
	case "powershell":
		return cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
	default:
		return fmt.Errorf("unsupported shell %q (bash|zsh|fish|powershell)", shell)
	}
}
