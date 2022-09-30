package foundation

type Bootstrapper interface {
	Bootstrap(*Application)
}

type Application struct {
	BasePath string
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
