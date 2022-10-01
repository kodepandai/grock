package bootstrap

import (
	"grock/src/foundation"
)

type LoadConfiguration struct{}

func (l LoadConfiguration) Bootstrap(app *foundation.Application) {
	configs := app.GetConfigs()

	// check and load Application config
	if appConfig, ok := configs["app"]; ok {
		app.Config = appConfig.(foundation.AppConfig)
	} else {
		panic("Cannot find app config")
	}

}
