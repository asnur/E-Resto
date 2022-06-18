package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"log"
	"testing"

	qrcode "github.com/skip2/go-qrcode"
)

func TestQRCode(t *testing.T) {
	code, err := qrcode.New("https://example.org", qrcode.Low)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	errs := png.Encode(&buf, code.Image(256))
	if errs != nil {
		log.Fatal(errs)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}
