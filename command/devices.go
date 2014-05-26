package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// DevicesCommand is a Command implementation that tells a spark cloud
// to list all devices associated with account identified by access token
type DevicesCommand struct {
	Ui cli.Ui
}

func (c *DevicesCommand) Help() string {
	helpText := `
Usage: spark devices

  Lists all devices associated with this account
`
	return strings.TrimSpace(helpText)
}

func (c *DevicesCommand) Run(args []string) int {

	client, err := AuthenticatedSparkClient(true)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to Spark cloud: %s", err))
		return 1
	}
	// defer client.Close()

	devices, _, err := client.Devices.List()

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error listing devices: %s", err))
		return 1
	}
	for _, device := range devices {
		c.Ui.Output(fmt.Sprintf("- %s", device))
	}

	return 0
}

func (c *DevicesCommand) Synopsis() string {
	return "Lists devices for authenticated user"
}
