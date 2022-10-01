package bootstrap

import "grock/src/foundation"

type RegisterProviders struct{}

func (r RegisterProviders) Bootstrap(app *foundation.Application) {
	app.RegisterConfiguredProviders()
}
