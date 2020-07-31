package database

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Config for DB connection
type Config struct {
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"MYSQL_DATABASE"`
	DbUser     string `env:"MYSQL_USER"`
	DbPassword string `env:"MYSQL_PASSWORD"`
}

func Connect() (*gorm.DB, error) {
	cfg := Config{}

	env.Parse(&cfg)

	dsn := "host="+ cfg.DbHost +" port=3306 user="+ cfg.DbUser +" dbname="+ cfg.DbName +" password="+ cfg.DbPassword +""

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}