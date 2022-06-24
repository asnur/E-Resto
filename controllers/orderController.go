package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func OrderDetail(c *fiber.Ctx) error {
	return c.Render("order", fiber.Map{
		"Title": "Order Detail",
	})
}

func Order(c *fiber.Ctx) error {
	no_meja := c.Params("id_meja")
	sess, err := store.Get(c)
	if err != nil {
		log.Println(err)
	}

	sess.Set("no_meja", no_meja)

	return c.Render("order", fiber.Map{
		"Title":   "Order",
		"No_Meja": sess.Get("no_meja"),
	})
}
