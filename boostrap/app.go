package boostrap

import (
	"grock/app/http"
	"grock/src/illuminate/foundation"
	"os"
)
var App *foundation.Application
func init() {
 // get current working directory, set as basePath
  cwd, err:= os.Getwd()
	if err !=nil {
	  panic(err)
	}
	App = foundation.CreateApplication(cwd)
	App.Singleton("HttpKernel", http.Kernel{})
}
