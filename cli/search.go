package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/rburmorrison/hoist/client"
	"github.com/spf13/cobra"
)

type searchOptions struct {
	Image string
}

// NewSearchCommand creates a command to search repos
// in a Docker registry.
func NewSearchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search REPO:TAG",
		Short: "Search for an image in a registry",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("search text must be provided")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSearch(searchOptions{
				Image: args[0],
			})
		},
	}

	return cmd
}

func runSearch(opts searchOptions) error {
	image, err := client.Search(opts.Image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(image)

	return nil
}
