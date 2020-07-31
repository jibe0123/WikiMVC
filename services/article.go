package services

import (
	"fmt"
	"github.com/jibe0123/WikiMVC/models"
	"log"
	"net/http"
)


func GetArticles() ([]models.Article, error) {
	articles, err := models.GetArticles()
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return []models.Article{}, err
	}

	return articles, nil
}


func GetArticle(article_id int) (models.Article, error) {
	article, err := models.GetArticle(article_id)
	if err != nil {
		log.Fatalf("Model execution: %s", err)
		return models.Article{}, err
	}

	return article, nil
}


func RenderArticles(w http.ResponseWriter) {
	articles, err := GetArticles()
	if err != nil {
		return
	}

	var render Render = Render{
		ParseFiles: []string{"web/articles.tmpl", "web/base.tmpl"},
		Writer:     w,
		Data:       articles,
	}

	err = RenderTemplate(render)
	if err != nil {
		fmt.Println(err.Error())
	}
}
