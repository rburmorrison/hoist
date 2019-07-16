package cliconfig

import (
	"fmt"
	"strings"

	"github.com/rburmorrison/hoist/client"
	"github.com/rburmorrison/hoist/types"
	"github.com/spf13/cobra"
)

type setOptions struct {
	Address string
	Mode    string
}

// NewSetCommand creates a command to set config
// options.
func NewSetCommand() *cobra.Command {
	var address string
	var mode string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set configuration options for hoist",
		Example: `hoist config set -a http://localhost:5000
hoist config set -a localhost:5000 -m http`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSet(setOptions{
				Address: address,
				Mode:    mode,
			})
		},
	}

	cmd.Flags().StringVarP(&address, "address", "a", "", "the address for the registry to explore")
	cmd.Flags().StringVarP(&mode, "mode", "m", "", "the method to send requests to the registry [http or https]")

	return cmd
}

func runSet(opts setOptions) error {
	// Address
	if opts.Address != "" {
		return setAddress(opts.Address)
	}

	// Mode
	if opts.Mode != "" {
		return setMode(opts.Mode)
	}

	return nil
}

func setAddress(address string) error {
	if err := client.ConfigSetAddress(address); err != nil {
		return err
	}

	if strings.HasPrefix(address, "http://") {
		address = strings.TrimPrefix(address, "http://")
		if err := client.ConfigSetMode(types.ModeHTTP); err != nil {
			return err
		}
		fmt.Println("Mode: HTTP")
	}

	if strings.HasPrefix(address, "https://") {
		address = strings.TrimPrefix(address, "https://")
		if err := client.ConfigSetMode(types.ModeHTTPS); err != nil {
			return err
		}
		fmt.Println("Mode: HTTPS")
	}

	fmt.Println("Address: " + address)
	return nil
}

func setMode(mode string) error {
	switch strings.ToUpper(mode) {
	case "HTTP":
		if err := client.ConfigSetMode(types.ModeHTTP); err != nil {
			return err
		}
		fmt.Println("Mode: HTTP")
		return nil
	case "HTTPS":
		if err := client.ConfigSetMode(types.ModeHTTPS); err != nil {
			return err
		}
		fmt.Println("Mode: HTTPS")
		return nil
	}

	return nil
}
