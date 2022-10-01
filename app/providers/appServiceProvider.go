package providers

import (
	"grock/src/foundation"
)

type AppServiceProvider struct {
	App *foundation.Application
}

func (s *AppServiceProvider) Register() {
}

func (s *AppServiceProvider) Boot() {
}

func (s *AppServiceProvider) Init(app *foundation.Application) {
	s.App = app
}
