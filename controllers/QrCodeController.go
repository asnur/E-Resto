package controllers

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/skip2/go-qrcode"
)

func GetQRCode(c *fiber.Ctx) error {
	env := godotenv.Load()
	if env != nil {
		log.Println("Error loading .env file")
	}
	id_meja := c.Params("id_meja")
	enc := CreateEncodeQR(os.Getenv("BASE_URL") + "/order/" + base64.RawStdEncoding.EncodeToString([]byte(id_meja)))

	return c.Render("qrcode", fiber.Map{
		"Data":  enc,
		"Title": "QR Code",
	})

}

func CreateEncodeQR(params string) string {
	code, err := qrcode.New(params, qrcode.Low)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	errs := png.Encode(&buf, code.Image(256))
	if errs != nil {
		log.Fatal(errs)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	return enc
}
