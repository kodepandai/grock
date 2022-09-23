package main

import (
	"fmt"
	"grock/app/http"
	"grock/boostrap"
	"os"
)

type User struct {
	Id   int
	Name string
}

func main() {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			// if (reflect.TypeOf(err) == reflect.TypeOf(bootstrap.HttpException{})){
			// 	fmt.Println("ini erre http ")
			// 	return
			// }
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()
	app := boostrap.App

	Kernel := app.Make("HttpKernel", nil).(http.Kernel)
	Kernel.Start()
	port := os.Getenv("PORT")
	fmt.Printf("server run in port: %s \n", port)
}
