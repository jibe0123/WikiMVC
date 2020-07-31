package models


import (
	"fmt"
	"github.com/jibe0123/WikiMVC/common"
	"github.com/jibe0123/WikiMVC/database"
	"github.com/jinzhu/gorm"

)

type Article struct {
	gorm.Model
	Title string
	Description string
}

func GetArticles() ([]Article, error) {
	var articles []Article

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return articles, err
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	err = db.Find(&articles).Error // find articles
	if err != nil {
		fmt.Println(err.Error())
		return articles, err
	}

	return articles, nil
}

func GetArticle(id int) (Article, error) {
	var article Article

	db, err := database.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return article, err
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	db.First(&article, id) // find article with id

	return article, nil
}

func (a Article) Serialize() common.JSON {
	return common.JSON{
		"id":         a.ID,
		"title":       a.Title,
	}
}