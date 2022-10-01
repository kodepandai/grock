package boostrap

import (
	"grock/config"
	"grock/src/foundation"
	"os"
)

var App *foundation.Application

func init() {
	// get current working directory, set as basePath
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	App = foundation.CreateApplication(cwd)
	App.RegisterConfigs(map[string]interface{}{
		"app": config.App,
	})
}
