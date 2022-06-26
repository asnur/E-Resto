package controllers

import (
	conf "eresto/config"

	"github.com/gofiber/fiber/v2"
)

func GetTable(c *fiber.Ctx) error {
	conf.Connect(".env")
	no_meja := c.Cookies("no_meja")

	if no_meja == "" {
		return c.Redirect("/order/1")
	}
	return c.Redirect("/order/" + no_meja)
}
