package cli

import (
	"fmt"
	"math"
	"strings"

	"github.com/rburmorrison/hoist/client"
	"github.com/rburmorrison/hoist/types"
	"github.com/spf13/cobra"
)

// NewRepsCommand creates a new command to retrieve
// repositories from a Docker registry.
func NewRepsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reps",
		Short: "List repositories from the registry",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runReps()
		},
	}

	return cmd
}

func runReps() error {
	summaries, err := client.Reps()
	if err != nil {
		return err
	}

	displayRepsTable(summaries, 3)

	return nil
}

func displayRepsTable(summaries []types.RepositorySummary, padding int) {
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
