package config

import (
	"grock/src/config"
	"grock/src/support"
)

var App config.App

func init() {
	App = config.App{
		AppName: support.Env("APP_NAME", "GROCK"),
	}
}
