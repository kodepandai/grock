package config

import (
	"grock/src/config"
)

var App config.App

func init() {
	App = config.App{
		AppName: "GROCK",
	}
}
