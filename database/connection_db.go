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

	dsn := cfg.DbUser + ":" + cfg.DbPassword + "@" + cfg.DbHost + "/" + cfg.
		DbName + "?parseTime=true&charset=utf8"

	db, err := gorm.Open(dsn)

	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}