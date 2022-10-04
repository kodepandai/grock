package bootstrap

import "grock/src/foundation"

type HandleException struct{}

func (h HandleException) Bootstrap(app *foundation.Application) {
	app.ExceptionHandler.Register(&app.ExceptionHandler)
}
