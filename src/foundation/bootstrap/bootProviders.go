package bootstrap

import "grock/src/foundation"

type BootProviders struct{}

func (b BootProviders) Bootstrap(app *foundation.Application) {
	app.Boot()
}
