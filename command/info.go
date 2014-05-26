package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// InfoCommand is a Command implementation which retrieves info
// about the coreId from spark cloud
type InfoCommand struct {
	Ui cli.Ui
}

func (c *InfoCommand) Help() string {
	helpText := `
Usage: spark info {coreId}

  Get basic information about the given Core,
  including the custom variables and functions it has exposed.
`
	return strings.TrimSpace(helpText)
}

func (c *InfoCommand) Run(args []string) int {
	coreId := args[0]

	client, err := AuthenticatedSparkClient(true)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed connecting to spark cloud: %s", err))
		return 1
	}

	info, _, err := client.Devices.Info(coreId)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed retrieving info: %s", err))
		return 1
	}
	c.Ui.Output(fmt.Sprintf("%v", info))
	return 0
}

func (c *InfoCommand) Synopsis() string {
	return "Displays basic information about the given Core"
}
