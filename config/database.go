package config

import (
	"log"
	"os"

	e "eresto/entities"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Query *gorm.DB

func Connect(env_location string) error {
	var err error

	env := godotenv.Load(env_location)
	if env != nil {
		log.Println("Error loading .env file")
	}
	Query, err = gorm.Open(mysql.Open(os.Getenv("DSN_MYSQL")), &gorm.Config{})

	if err != nil {
		return err
	}

	Query.AutoMigrate(&e.Table{}, &e.Order{}, &e.OrderDetail{}, &e.Menu{}, &e.Category{}, &e.Payement{})

	return nil
}
