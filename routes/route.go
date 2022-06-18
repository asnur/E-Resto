package routes

import (
	"log"

	"github.com/gofiber/template/html"

	c "eresto/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	enggine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: enggine,
	})
	app.Get("/", c.GetTable)
	app.Get("/qrcode/:id_meja", c.GetQRCode)

	log.Fatal(app.Listen(":3000"))
}
