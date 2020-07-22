package router

import (
	"github.com/gorilla/mux"
	"github.com/jibe0123/WikiMVC/controller"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		Name:        "index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controller.Index,
	},
}