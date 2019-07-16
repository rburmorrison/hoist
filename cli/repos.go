package cli

import (
	"fmt"
	"math"
	"strings"

	"github.com/rburmorrison/hoist/client"
	"github.com/rburmorrison/hoist/types"
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

func displayReposTable(summaries []types.RepositorySummary, padding int) {
	mostCharacters := 0
	for _, summary := range summaries {
		if len(summary.Name) > mostCharacters {
			mostCharacters = len(summary.Name)
		}
	}

	width := int(math.Max(float64(mostCharacters), 4))
	width += padding

	fmt.Printf("NAME%sTAG COUNT\n", strings.Repeat(" ", mostCharacters-4+padding))
	for _, summary := range summaries {
		separation := strings.Repeat(" ", mostCharacters-len(summary.Name)+padding)
		fmt.Printf("%s%s%d\n", summary.Name, separation, summary.TagCount)
	}
}
