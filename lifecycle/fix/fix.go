// fix package contains the logics of the fix operations, fix operations
// are in charge of adapting our source code to comply with newer versions
// of the CLI.
package fix

import (
	"context"
	"fmt"

	"github.com/wawandco/oxplugins/plugins"
)

//HelpText resturns the help Text of build function

var _ plugins.Command = (*Command)(nil)
var _ plugins.PluginReceiver = (*Command)(nil)

type Command struct {
	fixers []Fixer
}

func (c Command) Name() string {
	return "fix"
}

func (c Command) ParentName() string {
	return ""
}

func (c Command) HelpText() string {
	return "adapts the source code to comply with newer versions of the CLI"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	fmt.Println("Running [fix] command")

	//Run each of the fixers registered.
	for _, fixer := range c.fixers {
		fmt.Printf("Fixer: %v\n", fixer.Name())
	}

	return nil
}

func (c *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {
		if ptool, ok := plugin.(Fixer); ok {
			c.fixers = append(c.fixers, ptool)
		}
	}
}
