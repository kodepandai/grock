package routes

import (
	"grock/app/http/controllers"
	"grock/src/routing"
)

func Api(r *routing.Router) {
	r.Prefix("/api").Group(func() {
		r.Get("/", []interface{}{&controllers.WelcomeController{}, "index"})
		r.Prefix("/user").Group(func() {
			r.Get("/", []interface{}{&controllers.UserController{}, "listUser"})
		})
	})
}
