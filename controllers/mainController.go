package controllers

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

type CartItem struct {
	Name     string
	ID_Menu  int
	Quantity int
	Price    int
	Total    int
}

type MenuOrder struct {
	Menu         string
	Total_Prices int
	Quantity     int
}

type Ordered struct {
	No_Meja     int
	Payement    string
	Total_Price int
	Status      string
}

var store = session.New()
