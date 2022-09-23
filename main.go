package main

import (
	"fmt"
	"grock/boostrap"
	"illuminate/foundation/bootstrap"
	"illuminate/foundation/http"
	"os"
	"reflect"
)

type User struct {
	Id   int
	Name string
}

func main() {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			if (reflect.TypeOf(err) == reflect.TypeOf(bootstrap.HttpException{})){
				fmt.Println("ini erre http ")
				return
			}
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()
	app := boostrap.App()

	Kernel := app.Make("HttpKernel", nil).(http.Kernel)
	Kernel.Start()
}
