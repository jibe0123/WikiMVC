package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jibe0123/WikiMVC/third_party"
	"net/http"
)


func GetArticle(w http.ResponseWriter, r *http.Request) {

	muxVars := mux.Vars(r)

	strID := muxVars["id"]
	fmt.Println(strID)
	// render 1 article
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	third_party.RenderArticles(w)
}

func ArticleFormNew(w http.ResponseWriter, req *http.Request) {
	// Create article from from POST
}

func CreateArticle(w http.ResponseWriter, req *http.Request) {
	// Create article from from POST
}