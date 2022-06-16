package config

import (
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	err := Connect()
	if err != nil {
		log.Println("Database Not Connected")
	} else {
		log.Println("Database Connected")
	}
}
