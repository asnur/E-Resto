package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Query *gorm.DB

func Connect() error {
	var err error

	env := godotenv.Load("../.env")
	if env != nil {
		log.Println("Error loading .env file")
	}
	Query, err = gorm.Open(mysql.Open(os.Getenv("DSN_MYSQL")), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil
}
