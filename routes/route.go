package routes

import (
	"log"

	c "eresto/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	app := fiber.New()

	app.Get("/", c.GetTable)

	log.Fatal(app.Listen(":3000"))
}
