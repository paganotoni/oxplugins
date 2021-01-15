package version

import (
	"context"
	"fmt"

	"github.com/wawandco/oxplugins/plugins"
)

var (
	// The version of the CLI
	version = "0.0.1"
)

var (
	// Version is a Command
	_ plugins.Command = (*Version)(nil)
)

// Version command will print X version.
type Version struct{}

func (b Version) Name() string {
	return "version"
}

func (b Version) ParentName() string {
	return ""
}

func (b Version) HelpText() string {
	return "returns the current version of Oxpecker CLI"
}

// Run prints the version of the Oxpecker cli
func (b *Version) Run(ctx context.Context, root string, args []string) error {
	fmt.Printf("Oxpecker version %v\n", version)

	return nil
}
