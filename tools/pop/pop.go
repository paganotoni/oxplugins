package pop

import (
	"github.com/gobuffalo/packd"
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxplugins/tools/pop/migrate"
)

func Plugins(migrations packd.Box) []plugins.Plugin {
	pl := []plugins.Plugin{
		&Command{},
	}

	return append(pl, migrate.Plugins(migrations)...)
}