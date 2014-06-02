package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// ReadCommand is a Command implementation that reads the value of
// a variable exposed by the spark core
type ReadCommand struct {
	Ui cli.Ui
}

func (c *ReadCommand) Help() string {
	helpText := `
Usage: spark var {coreId} {variable}

  Request the current value of a variable exposed by the core, e.g.,
  GET /v1/devices/0123456789abcdef01234567/temperature
`
	return strings.TrimSpace(helpText)
}

func (c *ReadCommand) Run(args []string) int {
	if len(args) != 2 {
		c.Ui.Error("A coreId and a variable name must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	coreId := args[0]
	varName := args[1]

	client, err := AuthenticatedSparkClient(true)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed connecting to spark cloud: %s", err))
		return 1
	}

	variable, resp, err := client.Devices.Read(coreId, varName)
	if err != nil {
		if resp.StatusCode == 404 {
			c.Ui.Error(fmt.Sprintf("%s is not available on your spark core [%s]", varName, coreId))
			return 1
		}

		c.Ui.Error(fmt.Sprintf("Failed reading variable: %s", err))
		return 1
	}

	var res string
	switch v := variable.Result.(type) {
	case int:
		res = fmt.Sprintf("%d", v)
	case float64:
		res = fmt.Sprintf("%f", v)
	case string:
		res = fmt.Sprintf("%s", v)
	default:
		c.Ui.Error("Unknown result type")
		return 1
	}

	c.Ui.Output(fmt.Sprintf("%s = %s", varName, res))
	return 0
}

func (c *ReadCommand) Synopsis() string {
	return "Reads the value of variables exposed by spark core"
}
