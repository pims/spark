package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// ClaimCommand is a Command implementation that tells spark cloud
// to claim given core
type ClaimCommand struct {
	Ui cli.Ui
}

func (c *ClaimCommand) Help() string {
	helpText := `
Usage: spark claim {coreid}

  Claims the spark core
`
	return strings.TrimSpace(helpText)
}

func (c *ClaimCommand) Run(args []string) int {
	if len(args) != 1 {
		c.Ui.Error("A coreId must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	coreId := args[0]
	client, err := AuthenticatedSparkClient(true)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to Spark cloud: %s", err))
		return 1
	}
	// defer client.Close()

	_, claimErr := client.Devices.Claim(coreId)

	if claimErr != nil {
		c.Ui.Error(fmt.Sprintf("Error claiming coreId: %s", claimErr))
		return 1
	}

	return 0
}

func (c *ClaimCommand) Synopsis() string {
	return "Claims a spark core"
}
