package cli

import (
	"fmt"
	"strings"

	"github.com/rburmorrison/hoist/client"
	"github.com/rburmorrison/hoist/types"
	"github.com/spf13/cobra"
)

type configOptions struct {
	Address string
	Mode    string
}

// NewConfigCommand creates a new hoist command.
func NewConfigCommand() *cobra.Command {
	var address string
	var mode string

	cmd := &cobra.Command{
		Use:               "config",
		Short:             "Set configuration options for hoist",
		PersistentPreRunE: persistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConfig(configOptions{
				Address: address,
				Mode:    mode,
			})
		},
	}

	cmd.Flags().StringVarP(&address, "address", "a", "", "the address for the registry to explore")
	cmd.Flags().StringVarP(&mode, "mode", "m", "", "the method to send requests to the registry [http or https]")

	return cmd
}

func runConfig(opts configOptions) error {
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
