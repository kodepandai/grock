package providers

import (
	"grock/routes"
	"grock/src/foundation"
	"grock/src/routing"
)

type RouteServiceProvider struct {
	App *foundation.Application
}

func (s *RouteServiceProvider) Init(app *foundation.Application) {
	s.App = app
}

func (s *RouteServiceProvider) Register() {
	s.App.Router = &routing.Router{}

}

func (s *RouteServiceProvider) Boot() {
	routes.Api(s.App.Router)
}
