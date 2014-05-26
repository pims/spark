package command

import (
	"code.google.com/p/gopass"
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// TokensCommand is a Command implementation that list all oauth tokens
// from spark cloud
type TokensCommand struct {
	Ui cli.Ui
}

func (c *TokensCommand) Help() string {
	helpText := `
Usage: spark tokens me@example.com

  List all the access tokens for the given username
`
	return strings.TrimSpace(helpText)
}

func (c *TokensCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("A username must be specified.")
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
		c.Ui.Error(fmt.Sprintf("Error connecting to Spark cloud: %s", err))
		return 1
	}

	tokens, resp, err := client.Tokens.List(username, password)

	if err != nil {
		if resp.StatusCode == 401 {
			c.Ui.Error("HTTP 401. Wrong password maybe?")
			return 1
		}

		c.Ui.Error(fmt.Sprintf("Error list all tokens: %s", err))
		return 1
	}
	c.Ui.Output("OAuth tokens:")
	for _, token := range tokens {
		c.Ui.Info(fmt.Sprintf("- %s", token))
	}

	return 0
}

func (c *TokensCommand) Synopsis() string {
	return "List all access tokens"
}
