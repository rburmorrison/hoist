package cli

import (
	"github.com/rburmorrison/hoist/client"
	"github.com/spf13/cobra"
)

// NewReposCommand creates a new command to retrieve
// repositories from a Docker registry.
func NewReposCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "repos",
		Short:   "List repositories from the registry",
		Aliases: []string{"reps", "repositories"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRepos()
		},
	}

	return cmd
}

func runRepos() error {
	summaries, err := client.Repos()
	if err != nil {
		return err
	}

	displayReposTable(summaries, 3)

	return nil
}
