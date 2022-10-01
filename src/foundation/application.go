package foundation

type Bootstrapper interface {
	Bootstrap(*Application)
}

type AppConfig struct {
	AppName   string
	Providers []ServiceProvider
}

type Application struct {
	BasePath string
	Config   AppConfig
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

func (app *Application) RegisterConfiguredProviders() {
	for _, provider := range app.Config.Providers {
		provider.Init(app)
		provider.Register()
	}
}
