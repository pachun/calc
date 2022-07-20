package main

import (
	"os"

	"github.com/pachun/calc/cli"
)

func main() {
	cli.NewConsole(os.Stdout).Run(os.Args)
}
