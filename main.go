package main

import (
	"fmt"
	"os"

	"dhemery.com/panelgen/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
