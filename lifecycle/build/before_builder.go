package build

import (
	"context"

	"github.com/wawandco/oxplugins/plugins"
)

// BeforeBuilder interface allows to identify the things
// that will run before the build process has started.
type BeforeBuilder interface {
	plugins.Plugin
	RunBeforeBuild(context.Context, string, []string) error
}
