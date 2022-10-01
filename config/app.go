package config

import (
	"grock/app/providers"
	"grock/src/foundation"
	"grock/src/support"
)

var App foundation.AppConfig

func init() {
	App = foundation.AppConfig{
		AppName: support.Env("APP_NAME", "GROCK"),
		Providers: []foundation.ServiceProvider{
			&providers.AppServiceProvider{},
		},
	}
}
