package controllers

import (
	"log"

	conf "eresto/config"
	e "eresto/entities"

	"github.com/gofiber/fiber/v2"
)

func OrderDetail(c *fiber.Ctx) error {
	return c.Render("order", fiber.Map{
		"Title": "Order Detail",
	})
}

func Order(c *fiber.Ctx) error {
	conf.Connect(".env")
	no_meja := c.Params("id_meja")
	var category []e.Category
	if err := conf.Query.Find(&category).Error; err != nil {
		return err
	}
	sess, err := store.Get(c)
	if err != nil {
		log.Println(err)
	}

	sess.Set("no_meja", no_meja)

	return c.Render("order", fiber.Map{
		"Title":    "Order",
		"No_Meja":  sess.Get("no_meja"),
		"Kateogri": category,
	})
}
