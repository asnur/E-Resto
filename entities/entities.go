package entities

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	ID   int `gorm:"primary_key"`
	Name string
}
