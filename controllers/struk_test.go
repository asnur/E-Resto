package controllers

import (
	conf "eresto/config"
	"testing"
)

func TestStruk(t *testing.T) {
	conf.Connect("../.env")
}
