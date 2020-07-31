package controller

import (
	"github.com/jibe0123/WikiMVC/third_party"
	"net/http"
)


func GetArticle(w http.ResponseWriter, r *http.Request) {
	// render 1 article
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	third_party.RenderArticles(w)
}

func CreateArticle(w http.ResponseWriter, req *http.Request) {
	// Create article from from POST
}