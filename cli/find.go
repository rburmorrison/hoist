package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/rburmorrison/hoist/client"
	"github.com/spf13/cobra"
)

type findOptions struct {
	Image string
}

// NewFindCommand creates a command to search repos
// in a Docker registry.
func NewFindCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "find REPO:TAG",
		Short:   "Find an image in a registry",
		Aliases: []string{"search"},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("image name must be provided")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFind(findOptions{
				Image: args[0],
			})
		},
	}

	return cmd
}

func runFind(opts findOptions) error {
	image, err := client.Find(opts.Image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(image)

	return nil
}
