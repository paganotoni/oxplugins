package dev

import (
	"context"
	"fmt"
	"sync"

	"github.com/wawandco/oxplugins/plugins"
)

var _ plugins.Command = (*Command)(nil)

// dev is the dev command.
type Command struct {
	developers []Developer
}

func (d Command) Name() string {
	return "dev"
}

func (d Command) ParentName() string {
	return ""
}

//HelpText resturns the help Text of build function
func (d Command) HelpText() string {
	return "calls NPM or yarn to start webpack watching the assetst"
}

// Run calls NPM or yarn to start webpack watching the assets
// Also starts refresh listening for the changes in Go files.
func (d *Command) Run(ctx context.Context, root string, args []string) error {
	var wg sync.WaitGroup
	for _, tool := range d.developers {
		// Each of the tools runs in parallel
		wg.Add(1)
		go func(t Developer) {
			err := t.Develop(ctx, root)
			if err != nil {
				fmt.Println(err)
			}

			wg.Done()
		}(tool)
	}

	wg.Wait()
	return nil
}

func (d *Command) Receive(plugins []plugins.Plugin) {
	for _, tool := range plugins {
		ptool, ok := tool.(Developer)
		if !ok {
			continue
		}

		d.developers = append(d.developers, ptool)
	}
}
