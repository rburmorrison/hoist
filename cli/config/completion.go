package cliconfig

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCompletionCommand creates a command that
// generates bash completion for hoist.
func NewCompletionCommand() *cobra.Command {
	var zsh bool

	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate bash completion script",
		Example: `hoist config completion >> ~/.bash_completion
hoist config completion --zsh > ~/.oh-my-zsh/completions/_hoist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if zsh {
				return cmd.Root().GenZshCompletion(os.Stdout)
			}

			return cmd.Root().GenBashCompletion(os.Stdout)
		},
	}

	cmd.Flags().BoolVarP(&zsh, "zsh", "z", false, "generate zsh completion script instead")

	return cmd
}
