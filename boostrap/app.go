package boostrap

import (
	"illuminate/foundation"
	"os"
)

func App() foundation.Application {
 // get current working directory, set as basePath
  cwd, err:= os.Getwd()
	if err !=nil {
	  panic(err)
	}
	return foundation.NewApplication(cwd)

}
