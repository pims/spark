package command

import (
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
	"time"
)

// ExecCommand is a Command implementation which remotely executes a function
// on the core
type ExecCommand struct {
	Ui cli.Ui
}

func (c *ExecCommand) Help() string {
	helpText := `
Usage: spark exec --timeout=5 {coreId} {funcName} {args}

  Call a function exposed by the core, with arguments passed in request body,
  e.g., POST /v1/devices/0123456789abcdef01234567/brew

  options:
    --timeout=5  timeout in seconds
`
	return strings.TrimSpace(helpText)
}

func (c *ExecCommand) Run(args []string) int {

	var timeoutValue int
	cmdFlags := flag.NewFlagSet("exec", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.IntVar(&timeoutValue, "timeout", 10, "timeout value in seconds")

	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Error(fmt.Sprintf("%v", err))
		return 1
	}

	// Note that the flag package requires all flags to appear before positional arguments
	// (otherwise the flags will be interpreted as positional arguments).
	if len(args) != 3 {
		c.Ui.Error("A coreId a function name and its arguments must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	coreId := args[0]
	funcName := args[1]
	arguments := args[2]

	timeout := time.Duration(time.Duration(timeoutValue) * time.Second)
	client, err := AuthenticatedSparkClientWithTimeout(true, timeout)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed connecting to spark cloud: %s", err))
		return 1
	}

	exec, resp, err := client.Devices.Exec(coreId, funcName, arguments)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed executing function: %s", err))
		c.Ui.Error(fmt.Sprintf("Response: %s", resp))
		return 1
	}

	if exec.ReturnValue == 1 {
		c.Ui.Info(fmt.Sprintf("Command “%s” with arguments “%s” was successfully executed on core %s[%s].", funcName, arguments, exec.Name, exec.Id))
	} else {
		c.Ui.Error("Command was not successfully executed.")
	}

	return 0
}

func (c *ExecCommand) Synopsis() string {
	return "Calls a function exposed by spark core"
}
