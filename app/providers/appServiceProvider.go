package providers

import (
	"fmt"
	"grock/src/foundation"
)

type AppServiceProvider struct {
	App *foundation.Application
}

func (s *AppServiceProvider) Init(app *foundation.Application) {
	s.App = app
}

func (s *AppServiceProvider) Register() {
	fmt.Println("app registered")
}

func (s *AppServiceProvider) Boot() {
	fmt.Println("app booted")
}
