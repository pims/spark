package command

import (
	// "fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// ExecCommand is a Command implementation which remotely executes a function
// on the core
type ExecCommand struct {
	Ui cli.Ui
}

func (c *ExecCommand) Help() string {
	helpText := `
Usage: spark exec {coreId} {idk}

  Call a function exposed by the core, with arguments passed in request body,
  e.g., POST /v1/devices/0123456789abcdef01234567/brew
`
	return strings.TrimSpace(helpText)
}

func (c *ExecCommand) Run(args []string) int {
	if len(args) != 3 {
		c.Ui.Error("A coreId a function name and its arguments must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	c.Ui.Error("Executing remote functions is not fully implemented at the moment.")
	return 1

	// coreId := args[0]
	// funcName := args[1]
	// arguments := args[2]

	// client, err := AuthenticatedSparkClient(true)
	// if err != nil {
	// 	c.Ui.Error(fmt.Sprintf("Failed connecting to spark cloud: %s", err))
	// 	return 1
	// }

	// resp, err := client.Devices.Exec(coreId, funcName)
	// if err != nil {
	// 	c.Ui.Error(fmt.Sprintf("Failed executing function: %s", err))
	// 	return 1
	// }

	// fmt.Printf("Resp: %s", resp)
	// return 0
}

func (c *ExecCommand) Synopsis() string {
	return "Calls a function exposed by the core"
}
