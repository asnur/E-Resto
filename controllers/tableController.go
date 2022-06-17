package controllers

import (
	conf "eresto/config"
	e "eresto/entities"

	"github.com/gofiber/fiber/v2"
)

func GetTable(c *fiber.Ctx) error {
	conf.Connect(".env")
	var table []e.Table
	if err := conf.Query.Find(&table).Error; err != nil {
		return err
	}
	return c.JSON(table)
}
