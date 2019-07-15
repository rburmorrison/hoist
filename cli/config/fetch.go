package cliconfig

import (
	"fmt"

	"github.com/rburmorrison/hoist/client"
	"github.com/spf13/cobra"
)

type fetchOptions struct {
	ShowAddress bool
	ShowMode    bool
}

// NewFetchCommand creates a command to set config
// options.
func NewFetchCommand() *cobra.Command {
	var address bool
	var mode bool

	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch configuration options for hoist",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFetch(fetchOptions{
				ShowAddress: address,
				ShowMode:    mode,
			})
		},
	}

	cmd.Flags().BoolVarP(&address, "address", "a", false, "display the address")
	cmd.Flags().BoolVarP(&mode, "mode", "m", false, "display the mode")

	return cmd
}

func runFetch(opts fetchOptions) error {
	// Address
	if opts.ShowAddress {
		address, err := client.ConfigFetchAddress()
		if err != nil {
			return err
		}
		fmt.Println("Address: " + address)
	}

	// Mode
	if opts.ShowMode {
		mode, err := client.ConfigFetchMode()
		if err != nil {
			return err
		}
		fmt.Println("Mode: " + mode.String())
	}

	return nil
}
