package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rburmorrison/hoist/client"
	"github.com/spf13/cobra"
)

type tagsOptions struct {
	Repository string
	Separator  string
}

// NewTagsCommand creates a new command to retrieve a
// repository's tags from a Docker registry.
func NewTagsCommand() *cobra.Command {
	var separator string

	cmd := &cobra.Command{
		Use:   "tags REPOSITORY",
		Short: "List repositories from the registry",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("a repository must be provided")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTags(tagsOptions{
				Repository: args[0],
				Separator:  separator,
			})
		},
	}

	cmd.Flags().StringVarP(&separator, "separator", "s", "\n", "string to separate tags with")

	return cmd
}

func runTags(opts tagsOptions) error {
	tags, err := client.Tags(opts.Repository)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		fmt.Printf("%s%s", tag, opts.Separator)
	}

	if !strings.HasSuffix(opts.Separator, "\n") {
		fmt.Println()
	}

	return nil
}
