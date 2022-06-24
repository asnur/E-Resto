package controllers

import (
	conf "eresto/config"
	e "eresto/entities"

	"github.com/gofiber/fiber/v2"
)

func CategoryMenu(c *fiber.Ctx) error {
	var category []e.Category
	if err := conf.Query.Find(&category).Error; err != nil {
		return err
	}
	return c.Render("category", fiber.Map{
		"Title": "Category",
		"Data":  category,
	})
}
