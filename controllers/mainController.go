package controllers

import "github.com/gofiber/fiber/v2/middleware/session"

type CartItem struct {
	Name     string
	ID_Menu  int
	Quantity int
	Price    int
	Total    int
}

var store = session.New()
