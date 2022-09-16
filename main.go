package main 

import (
	"fmt"
	"grock/boostrap"
)

func main() {
	app:= boostrap.App()
	app.Instances["x"] = "Instance X"
  fmt.Println(app)
}
