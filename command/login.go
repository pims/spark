package command

import (
	"code.google.com/p/gopass"
	"fmt"
	"github.com/mitchellh/cli"
	"io/ioutil"
	"strings"
)

// LoginCommand is a Command implementation that attempts to
// log in to spark cloud
type LoginCommand struct {
	Ui cli.Ui
}

func (c *LoginCommand) Help() string {
	helpText := `
Usage: spark login me@example.com

  Retrieves a token from spark cloud and stores it locally
`
	return strings.TrimSpace(helpText)
}

func (c *LoginCommand) Run(args []string) int {
	if len(args) != 2 {
		c.Ui.Error("A username and a password name must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	username := args[0]
	password, err := gopass.GetPass("Password: ")

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error reading password from prompt: %s", err))
		return 1
	}

	client, err := AuthenticatedSparkClient(false)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to spark cloud: %s", err))
		return 1
	}
	// defer client.Close()

	token, _, err := client.Tokens.Login(username, password)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error retrieving token: %s", err))
		return 1
	}

	err = ioutil.WriteFile(SettingsFileName, []byte(*token.Value), 0755)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error persisting token to file: %s", err))
		return 1
	}

	c.Ui.Output(fmt.Sprintf("Token = %v", token))
	return 0
}

func (c *LoginCommand) Synopsis() string {
	return "Log in spark cloud"
}
