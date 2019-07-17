package cliconfig

import "github.com/spf13/cobra"

// NewConfigCommand creates a group of commands for
// configuration.
func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage hoist configuration",
	}

	cmd.AddCommand(
		NewSetCommand(),
		NewFetchCommand(),
		NewCompletionCommand(),
	)

	return cmd
}
