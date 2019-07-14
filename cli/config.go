package cli

import (
	"fmt"

	"github.com/rburmorrison/hoist/client"
	"github.com/spf13/cobra"
)

type configOptions struct {
	Address string
}

// NewConfigCommand creates a new hoist command.
func NewConfigCommand() *cobra.Command {
	var address string

	cmd := &cobra.Command{
		Use:               "config",
		Short:             "Set configuration options for hoist",
		PersistentPreRunE: persistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConfig(configOptions{
				Address: address,
			})
		},
	}

	cmd.Flags().StringVarP(&address, "address", "a", "", "the address for the registry to explore")

	return cmd
}

func runConfig(opts configOptions) error {
	// Address
	if opts.Address != "" {
		if err := client.ConfigSetAddress(opts.Address); err != nil {
			return err
		}

		fmt.Println(opts.Address)
		return nil
	}

	return nil
}
