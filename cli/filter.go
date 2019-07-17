package cli

import (
	"errors"

	"github.com/rburmorrison/hoist/client"
	"github.com/spf13/cobra"
)

type filterOptions struct {
	Text string
}

// NewFilterCommand creates a command to filter repos
// in a Docker registry.
func NewFilterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "filter TEXT",
		Short: "Filter repos that contain some text",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("filter text must be provided")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFilter(filterOptions{
				Text: args[0],
			})
		},
	}

	return cmd
}

func runFilter(opts filterOptions) error {
	summaries, err := client.Filter(opts.Text)
	if err != nil {
		return err
	}

	displayReposTable(summaries, 3)

	return nil
}
