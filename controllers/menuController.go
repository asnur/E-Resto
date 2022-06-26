package controllers

import (
	conf "eresto/config"
	e "eresto/entities"

	"github.com/gofiber/fiber/v2"
)

func ListMenu(c *fiber.Ctx) error {
	conf.Connect(".env")
	var menu []e.Menu
	var payement []e.Payement
	id_menu := c.Params("id_menu")
	if err_menu := conf.Query.Where("id_category", id_menu).Find(&menu).Error; err_menu != nil {
		return err_menu
	}

	if err_payement := conf.Query.Find(&payement).Error; err_payement != nil {
		return err_payement
	}

	return c.Render("menu", fiber.Map{
		"Title":    "Menu",
		"Menu":     menu,
		"Payement": payement,
	})
}
