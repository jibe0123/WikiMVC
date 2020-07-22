package controller

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	PageTitle string
}


func Index(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

	data := PageData{
		PageTitle: "Index Page",
	}
	var err = tmpl.Execute(w, data)

	if err != nil {
		log.Print(err)
		return
	}
}