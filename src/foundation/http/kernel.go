package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"grock/src/foundation"
	"grock/src/foundation/bootstrap"
	grockHttp "grock/src/http"
	"grock/src/http/exception"
	"grock/src/http/response"
	"grock/src/support"
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"
)

type Kernel struct {
	App *foundation.Application

	Bootstrappers []foundation.Bootstrapper
}

/* Register default bootstrappers to Kernel */
func (k *Kernel) InitDefaultBootstrapper() {
	k.Bootstrappers = append(k.Bootstrappers,
		bootstrap.LoadEnvironmentVariable{},
		bootstrap.LoadConfiguration{},
		bootstrap.HandleException{},
		bootstrap.RegisterProviders{},
		bootstrap.BootProviders{},
	)
	// TODO: register more bootstrappers here
}

/* Start Kernel */
func (k *Kernel) Start() {
	k.InitDefaultBootstrapper()
	k.App.BootstrapWith(k.Bootstrappers)

	server := httprouter.New()
	server.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(exception.NotFoundHttpException("page not found"))
	})

	server.PanicHandler = func(w http.ResponseWriter, r *http.Request, e interface{}) {
		response := k.renderException(e)
		k.Send(w, response)
	}

	routes := k.App.Router.GetRoutes()
	for _, route := range routes {
		uri := route.Uri
		action := route.Action
		method := route.Method
		server.Handle(method, uri, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			actionResult := action.Call(nil)
			var res grockHttp.Response
			switch val := actionResult[0].Interface().(type) {
			case grockHttp.Response:
				res = val
			case grockHttp.Responsable:
				res = val.ToResponse()
			default:
				res = response.Make(val, http.StatusOK, nil)
			}
			k.Send(w, res)
		})
	}
	port := support.Env("PORT", "8000")
	done := make(chan bool)

	// run server in separate go routine
	go http.ListenAndServe(":"+port, server)

	fmt.Printf("server run on http://localhost:%s \n", port)

	<-done
}

/* Render Exception to client */
func (k *Kernel) renderException(e interface{}) grockHttp.Response {
	switch v := e.(type) {
	case string:
		e = errors.New(v)
	case error:
		break
	default:
		e = exception.HttpException{Message: "Server Error", StatusCode: http.StatusInternalServerError}
	}
	return k.App.ExceptionHandler.Render(e.(error))
}

/* Send formated http response to client */
func (k *Kernel) Send(w http.ResponseWriter, r grockHttp.Response) {
	w.WriteHeader(r.Status)
	if reflect.TypeOf(r).Kind() == reflect.Struct {
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(r.Original)
		if err != nil {
			panic("Cannot convert response to json")
		}
		w.Write(jsonResp)
		return
	}
	if reflect.TypeOf(r).Kind() == reflect.String {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(r.Original.(string)))
	}
}
