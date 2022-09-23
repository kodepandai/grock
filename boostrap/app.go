package boostrap

import (
	"illuminate/foundation"
	"illuminate/foundation/http"
	"os"
)

func App() foundation.Application {
 // get current working directory, set as basePath
  cwd, err:= os.Getwd()
	if err !=nil {
	  panic(err)
	}
	app:= foundation.CreateApplication(cwd)
	app.Singleton("HttpKernel", http.Kernel{App:app})

	return app
}
