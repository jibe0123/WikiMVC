package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)


func Connect() (*gorm.DB, error) {
	user := os.Getenv("MYSQL_USER")
	pwd := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("DB_HOST")

	dsn := user + ":" + pwd + "@tcp("+ host +")/" + database + "?parseTime=true&charset=utf8"

	var err error

	for i:=1; i <= 3; i++ {
		db, err := gorm.Open("mysql", dsn)

		if err == nil {
			fmt.Println("Connected to database")
			return db,nil
		}
		fmt.Println(err)
		time.Sleep(10 * time.Second)
	}

	return nil, err
}