package controller

import (
	"github.com/jibe0123/WikiMVC/models"
	"github.com/jibe0123/WikiMVC/strategies"
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


func ExportCsv(w http.ResponseWriter, r *http.Request) {

	var activeStrategy strategies.ExportStrategy

	var article models.Article

	activeStrategy = &strategies.CsvFile{Name: "file.csv"}

	err := activeStrategy.Export(article)

	if err != nil {
		log.Fatal(err)
	}
}
