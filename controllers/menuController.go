package controllers

import (
	conf "eresto/config"
	e "eresto/entities"

	"github.com/gofiber/fiber/v2"
)

func ListMenu(c *fiber.Ctx) error {
	conf.Connect(".env")
	var menu []e.Menu
	id_menu := c.Params("id_menu")
	if err := conf.Query.Where("id_category", id_menu).Find(&menu).Error; err != nil {
		return err
	}

	return c.Render("menu", fiber.Map{
		"Title": "Menu",
		"Menu":  menu,
	})
}
