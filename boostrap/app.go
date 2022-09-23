package boostrap

import (
	"grock/src/illuminate/foundation"
	"grock/app/http"
	"os"
)

func App() foundation.Application {
 // get current working directory, set as basePath
  cwd, err:= os.Getwd()
	if err !=nil {
	  panic(err)
	}
	app:= foundation.CreateApplication(cwd)
	app.Singleton("HttpKernel", http.Kernel{})

	return app
}
