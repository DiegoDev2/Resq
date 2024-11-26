package main

import (
	"fmt"
	"os"

	"github.com/DiegoDev2/resq/cli"
	"github.com/fatih/color"
)

const banner = `
 ______ _____ _____  _____
| ___ \  ___/  ___||  _  |
| |_/ / |__ \ ` + "`" + `--. | | | |
|    /|  __| ` + "`" + `--. \| | | |
| |\ \| |___/\__/ /\ \/' /
\_| \_\____/\____/  \_/\_\

`

func main() {
	color.Green(banner)
	if err := cli.Command().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
