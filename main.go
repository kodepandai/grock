package main 

import (
	"fmt"
	"grock/boostrap"
)

type User struct {
	Id int
	Name string
}
func main() {
	app:= boostrap.App()

	x:=app.Make("HttpKernel", nil).(func()string)
	fmt.Println(x())
}

