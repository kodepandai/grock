package foundation

import (
	"grock/src/illuminate/container"
)

type Bootstrapper interface {
  Bootstrap(*Application)
}

type Application struct {
  container.Container
  BasePath string
}

func CreateApplication(basePath string) *Application {
  app:= Application {BasePath: basePath}
  app.Container.Init()
  return &app
}

func(app *Application) BootstrapWith(bootstrappers []Bootstrapper){
	for _, bootstrapper := range bootstrappers {
		bootstrapper.Bootstrap(app)
	}
}
