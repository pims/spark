package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"io/ioutil"
	"strings"
)

// LogoutCommand is a Command implementation that attempts to
// delete the access token stored locally
type LogoutCommand struct {
	Ui cli.Ui
}

func (c *LogoutCommand) Help() string {
	helpText := `
Usage: spark logout

  Empties local settings file which contains the most recent access token.
`
	return strings.TrimSpace(helpText)
}

func (c *LogoutCommand) Run(args []string) int {
	if len(args) > 0 {
		c.Ui.Error("This command does not accept any argument.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	err := ioutil.WriteFile(SettingsFileName, []byte(""), 0755)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error removing token from file: %s", err))
		return 1
	}

	c.Ui.Info("Successfully logged out")
	return 0
}

func (c *LogoutCommand) Synopsis() string {
	return "Logout from spark cloud"
}
