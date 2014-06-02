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
	if len(args) != 1 {
		c.Ui.Error("A coreId must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	coreId := args[0]

	client, err := AuthenticatedSparkClient(true)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed connecting to spark cloud: %s", err))
		return 1
	}

	info, resp, err := client.Devices.Info(coreId)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed retrieving info: %s", err))
		c.Ui.Error(fmt.Sprintf("Response: %s", resp))
		return 1
	}

	c.Ui.Info(fmt.Sprintf("Info for %s [%s]", info.Name, info.Id))
	c.Ui.Info("Variables:")
	for varName, varType := range info.Variables {
		c.Ui.Info(fmt.Sprintf(" - %s : %s", varName, varType))
	}

	c.Ui.Info("Functions:")
	for _, funcName := range info.Functions {
		c.Ui.Info(fmt.Sprintf(" - %s", funcName))
	}
	return 0
}

func (c *InfoCommand) Synopsis() string {
	return "Displays basic information about the given Core"
}
