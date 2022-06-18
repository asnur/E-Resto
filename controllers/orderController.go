package controllers

import "github.com/gofiber/fiber/v2"

func OrderDetail(c *fiber.Ctx) error {
	return c.Render("order", fiber.Map{
		"Title": "Order Detail",
	})
}
