package controller

import (
	"github.com/jibe0123/WikiMVC/services"
	"net/http"
)


func GetArticles(w http.ResponseWriter, r *http.Request) {
	services.RenderArticles(w)
}
