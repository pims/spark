package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
)

const (
	SettingsFileName = ".sparkio"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	args := os.Args[1:]

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("spark"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
