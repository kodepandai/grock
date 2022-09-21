package foundation

import "illuminate/container"

type Application struct {
  container.Container
  BasePath string
}

func CreateApplication(basePath string) Application {
  app:= Application {BasePath: basePath}
  app.Container.Init()
  return app
}

