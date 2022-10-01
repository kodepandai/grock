package bootstrap

import (
	"fmt"
	"grock/src/config"
	"grock/src/foundation"
)

type LoadConfiguration struct{}

func (l LoadConfiguration) Bootstrap(app *foundation.Application) {
	configs := app.GetConfigs()

	// check and load Application config
	if appConfig, ok := configs["app"]; ok {
		app.Config = appConfig.(config.App)
	} else {
		panic("Cannot find app config")
	}
	fmt.Println(app.Config)

}
