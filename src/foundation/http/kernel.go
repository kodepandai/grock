package http

import (
	"grock/src/foundation"
	"grock/src/foundation/bootstrap"
)

type Kernel struct {
	App *foundation.Application

	Bootstrappers []foundation.Bootstrapper
}

/* Register defaust bootstrappers to Kernel */
func (k *Kernel) InitDefaultBootstrapper() {
	k.Bootstrappers = append(k.Bootstrappers,
		bootstrap.LoadEnvironmentVariable{},
		bootstrap.LoadConfiguration{},
		bootstrap.RegisterProviders{},
	)
	// TODO: register more bootstrappers here
}

/* Start Kernel */
func (k *Kernel) Start() {
	k.InitDefaultBootstrapper()
	k.App.BootstrapWith(k.Bootstrappers)
}
