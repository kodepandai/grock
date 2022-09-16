package foundation

import "illuminate/container"

type Application struct {
  container.Container
  BasePath string
}

func NewApplication(basePath string) Application {
  app:= Application {BasePath: basePath}
  app.initData()
  return app
}

func (app *Application) initData(){
  app.Instances = make(map[string]interface{})
}

