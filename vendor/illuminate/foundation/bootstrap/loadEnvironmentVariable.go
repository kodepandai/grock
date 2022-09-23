package bootstrap

import (
 "illuminate/foundation"
 "fmt"
)

type HttpException struct {
  message string
}
type LoadEnvironmentVariable struct {}

func (l LoadEnvironmentVariable) Bootstrap(app *foundation.Application){
  fmt.Println("ENVVVV")
  panic(HttpException{message: "Env error"})
}
