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
	Route{
		Name:        "all_articles",
		Method:      "GET",
		Pattern:     "/articles",
		HandlerFunc: controller.GetArticles,
	},
	Route{
		Name:        "article",
		Method:      "GET",
		Pattern:     "/article/{id}",
		HandlerFunc: controller.GetArticle,
	},
	Route{
		Name:        "new_article",
		Method:      "POST",
		Pattern:     "/article/new",
		HandlerFunc: controller.ArticleFormNew,
	},
	Route{
		Name:        "new_article",
		Method:      "GET",
		Pattern:     "/article/new",
		HandlerFunc: controller.CreateArticle,
	},
	Route{
		Name:        "export csv",
		Method:      "GET",
		Pattern:     "/export-csv",
		HandlerFunc: controller.ExportCsv,
	},
}