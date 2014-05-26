package main

import (
	"github.com/mitchellh/cli"
	"github.com/pims/spark/command"
	"os"
)

// Commands is the mapping of all the available Spark commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.ColoredUi{
		InfoColor:  cli.UiColorGreen,
		ErrorColor: cli.UiColorRed,
		Ui:         &cli.BasicUi{Writer: os.Stdout},
	}

	Commands = map[string]cli.CommandFactory{
		"login": func() (cli.Command, error) {
			return &command.LoginCommand{
				Ui: ui,
			}, nil
		},

		"rename": func() (cli.Command, error) {
			return &command.RenameCommand{
				Ui: ui,
			}, nil
		},

		"devices": func() (cli.Command, error) {
			return &command.DevicesCommand{
				Ui: ui,
			}, nil
		},

		"claim": func() (cli.Command, error) {
			return &command.ClaimCommand{
				Ui: ui,
			}, nil
		},

		"info": func() (cli.Command, error) {
			return &command.InfoCommand{
				Ui: ui,
			}, nil
		},

		"read": func() (cli.Command, error) {
			return &command.ReadCommand{
				Ui: ui,
			}, nil
		},

		"exec": func() (cli.Command, error) {
			return &command.ExecCommand{
				Ui: ui,
			}, nil
		},
		"tokens": func() (cli.Command, error) {
			return &command.TokensCommand{
				Ui: ui,
			}, nil
		},
	}
}
