package http

import (
	"grock/src/illuminate/foundation"
	"grock/src/illuminate/foundation/bootstrap"
)

type Kernel struct {
  App foundation.Application

  Bootstrappers []foundation.Bootstrapper
}

func(k *Kernel) InitDefaultBootstrapper(){
  k.Bootstrappers = append(k.Bootstrappers, bootstrap.LoadEnvironmentVariable{})
}


func(k *Kernel) Start(){
  k.InitDefaultBootstrapper()
	k.App.BootstrapWith(k.Bootstrappers)
}