package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelfcads/file-api/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"DocumentCreate", "POST", "/documents", handler.Document},
	Route{"DocumentGet", "GET", "/documents/{key}", handler.Get},
	Route{"Healthcheck", "GET", "/healthcheck", handler.Healthcheck},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1").Subrouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
