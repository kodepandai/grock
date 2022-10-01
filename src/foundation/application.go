package foundation

import "grock/src/config"

type Bootstrapper interface {
	Bootstrap(*Application)
}

type Application struct {
	BasePath string
	Config   config.App
	configs  map[string]interface{}
}

func CreateApplication(basePath string) *Application {
	app := Application{BasePath: basePath}
	return &app
}

func (app *Application) BootstrapWith(bootstrappers []Bootstrapper) {
	for _, bootstrapper := range bootstrappers {
		bootstrapper.Bootstrap(app)
	}
}

func (app *Application) RegisterConfigs(configs map[string]interface{}) {
	app.configs = configs
}

func (app *Application) GetConfigs() map[string]interface{} {
	return app.configs
}
