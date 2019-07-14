package cli

import (
	"github.com/spf13/cobra"
)

// NewHoistCommand creates a new hoist command.
func NewHoistCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "hoist",
		Short:             "A CLI to explore private Docker Registries",
		PersistentPreRunE: persistentPreRunE,
	}

	cmd.AddCommand(
		NewConfigCommand(),
	)

	return cmd
}
