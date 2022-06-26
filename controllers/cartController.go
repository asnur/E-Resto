package controllers

import (
	conf "eresto/config"
	e "eresto/entities"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

var cart []CartItem

func Cart(c *fiber.Ctx) error {
	id_menu, _ := strconv.Atoi(c.FormValue("id_menu"))
	quantity, _ := strconv.Atoi(c.FormValue("quantity"))
	sess, err := store.Get(c)
	if err != nil {
		log.Println(err)
	}

	//check index array
	idx := slices.IndexFunc(cart, func(c CartItem) bool { return c.ID_Menu == id_menu })

	if idx == -1 {
		//append to cart
		item, _ := GetItem(id_menu)
		cart = append(cart, CartItem{Name: item.Name, ID_Menu: id_menu, Quantity: quantity, Price: item.Price, Total: item.Price * quantity})
	} else {
		//update quantity
		cart[idx].Quantity += quantity
		cart[idx].Total = cart[idx].Price * cart[idx].Quantity
	}

	sess.Set("cart", cart)

	return c.JSON(fiber.Map{"cart": cart})
}

func GetItem(id_menu int) (e.Menu, error) {
	conf.Connect(".env")
	var menu e.Menu
	if err := conf.Query.Where("id = ?", id_menu).First(&menu).Error; err != nil {
		return menu, err
	}
	return menu, nil
}
