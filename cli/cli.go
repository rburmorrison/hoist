// Package cli defines the cobra commands used for
// the hoist command-line tool.
package cli

import (
	"os"

	"github.com/rburmorrison/hoist/types"
	"github.com/spf13/cobra"
)

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(types.ConfigPath); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(types.ConfigPath, 770); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
