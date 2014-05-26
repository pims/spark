package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// RenameCommand is a Command implementation that tells spark cloud
// to rename a given core
type RenameCommand struct {
	Ui cli.Ui
}

func (c *RenameCommand) Help() string {
	helpText := `
Usage: spark rename {coreid} {new_name}

  Renames {coreid} to {new_name}.
  Ex: spark rename 123456789 my-awesome-core
`
	return strings.TrimSpace(helpText)
}

func (c *RenameCommand) Run(args []string) int {
	if len(args) != 2 {
		c.Ui.Error("A coreId and a name must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	coreId := args[0]
	name := args[1]

	client, err := AuthenticatedSparkClient(true)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to Spark client: %s", err))
		return 1
	}
	// defer client.Close()

	resp, err := client.Devices.Rename(coreId, name)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Got HTTP %d", resp.StatusCode))
		c.Ui.Error(fmt.Sprintf("Error renaming coreId: %s", err))
		return 1
	}

	c.Ui.Info(fmt.Sprintf("Successfully renamed core %s to %s", coreId, name))
	return 0
}

func (c *RenameCommand) Synopsis() string {
	return "Renames a core"
}
