package command

import (
	"code.google.com/p/gopass"
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// DeleteCommand is a Command implementation that deletes a given
// access token from spark cloud
type InvalidateCommand struct {
	Ui cli.Ui
}

func (c *InvalidateCommand) Help() string {
	helpText := `
Usage: spark invalidate 6fe38619bc260da99e3528f6b22d2fa5c35d0d57

  Invalidates the given access token
`
	return strings.TrimSpace(helpText)
}

func (c *InvalidateCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("An access token must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	token := args[0]
	username, err := c.Ui.Ask("Username: ")

	if err != nil {
		c.Ui.Error("Failed reading username from prompt.")
		return 1
	}

	password, err := gopass.GetPass("Password: ")

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error reading password from prompt: %s", err))
		return 1
	}

	client, err := AuthenticatedSparkClient(false)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to Spark cloud: %s", err))
		return 1
	}

	resp, err := client.Tokens.Delete(token, username, password)

	if err != nil {
		if resp != nil && resp.StatusCode == 401 {
			c.Ui.Error("HTTP 401. Wrong password maybe?")
			return 1
		}

		c.Ui.Error(fmt.Sprintf("Error invalidate token: %s", err))
		return 1
	}

	c.Ui.Info("Access token successfully invalidated.")
	return 0
}

func (c *InvalidateCommand) Synopsis() string {
	return "Invalidates an access token. Requires username/password"
}
