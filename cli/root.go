package cli

import (
	cliconfig "github.com/rburmorrison/hoist/cli/config"
	"github.com/spf13/cobra"
)

// NewHoistCommand creates a new hoist command.
func NewHoistCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "hoist",
		Short:             "A CLI to explore private Docker Registries",
		Version:           "v1.0.1",
		PersistentPreRunE: persistentPreRunE,
	}

	cmd.AddCommand(
		NewReposCommand(),
		NewTagsCommand(),
		cliconfig.NewConfigCommand(),
	)

	return cmd
}
