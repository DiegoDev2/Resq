package main

import (
	"fmt"
	"os"

	"github.com/DiegoDev2/resq/cli"
)

func main() {
	if err := cli.Command().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
