package main

import (
	"fmt"
	"os"

	"github.com/kierandotai/defillama-pp-cli/internal/cli"
)

var version = "0.1.0"

func main() {
	root := cli.New(version)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
