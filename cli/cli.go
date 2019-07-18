// Package cli defines the cobra commands used for
// the hoist command-line tool.
package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/rburmorrison/hoist/types"
	"github.com/spf13/cobra"
)

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	// Create the configuration directory if it does not
	// exist yet
	if _, err := os.Stat(types.ConfigDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(types.ConfigDir, 0770); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Create an empty configuration file if it does not
	// exist yet
	if _, err := os.Stat(types.ConfigFilePath); err != nil {
		if os.IsNotExist(err) {
			// Define defaults
			config := make(types.Configuration)
			config["address"] = "localhost:5000"
			config["mode"] = types.ModeHTTP

			// Generate empty JSON data
			bs, err := json.Marshal(config)
			if err != nil {
				return err
			}

			// Write the data to the file
			if err = ioutil.WriteFile(types.ConfigFilePath, append(bs, '\n'), 0770); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func displayReposTable(summaries []types.RepositorySummary, padding int) {
	mostCharacters := 0
	for _, summary := range summaries {
		if len(summary.Name) > mostCharacters {
			mostCharacters = len(summary.Name)
		}
	}

	width := int(math.Max(float64(mostCharacters), 4))
	width += padding

	fmt.Printf("NAME%sTAG COUNT\n", strings.Repeat(" ", width-4+padding))
	for _, summary := range summaries {
		separation := strings.Repeat(" ", width-len(summary.Name)+padding)
		fmt.Printf("%s%s%d\n", summary.Name, separation, summary.TagCount)
	}
}
