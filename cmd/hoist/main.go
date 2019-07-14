package main

import (
	"fmt"
	"os"

	"github.com/rburmorrison/hoist/cli"
)

func main() {
	if err := cli.NewHoistCommand().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
}
