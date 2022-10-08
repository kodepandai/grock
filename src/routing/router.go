package routing

import (
	"grock/src/http"
	"reflect"
	"strings"
)

type RouteData struct {
	Prefix string
	Uri    string
	Method string
	Action reflect.Value
}

type Router struct {
	routes      []RouteData
	prefixStack []string
	deep        int
}

const (
	GET  string = "GET"
	POST        = "POST"
)

func (r *Router) AddRoute(method string, path string, routerAction []interface{}) *Router {
	prefix := strings.Join(r.prefixStack, "")
	controller := routerAction[0].(http.Controller)
	controller.Init()
	controllerMethod := routerAction[1].(string)
	action := reflect.ValueOf(controller).MethodByName(strings.Title(controllerMethod))
	if !action.IsValid() {
		panic(reflect.ValueOf(controller).String() + " doesn't have method " + controllerMethod)
	}
	r.routes = append(r.routes, RouteData{
		Prefix: prefix,
		Uri:    prefix + path,
		Method: method,
		Action: action,
	})
	return r
}

func (r *Router) Prefix(prefix string) *Router {
	r.prefixStack = append(r.prefixStack, prefix)
	return r
}

func (r *Router) Group(callback func()) {
	r.deep++
	callback()
	r.prefixStack = r.prefixStack[:len(r.prefixStack)-1]
	r.deep--
}

func (r *Router) Get(uri string, action []interface{}) {
	r.AddRoute(GET, uri, action)
}

func (r *Router) GetRoutes() []RouteData {
	return r.routes
}
