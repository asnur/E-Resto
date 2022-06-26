package controllers

import (
	"strconv"
	"time"

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
	cookie := new(fiber.Cookie)
	cookie.Name = "no_meja"
	cookie.Value = no_meja
	cookie.Expires = time.Now().Add(time.Hour * 24)
	c.Cookie(cookie)

	return c.Render("order", fiber.Map{
		"Title":    "Order",
		"No_Meja":  no_meja,
		"Kateogri": category,
	})
}

func CheckOut(c *fiber.Ctx) error {
	conf.Connect(".env")
	no_meja, _ := strconv.Atoi(c.Cookies("no_meja"))
	total, _ := strconv.Atoi(c.FormValue("total"))
	payment, _ := strconv.Atoi(c.FormValue("payment"))
	order := e.Order{TableNo_Meja: no_meja, Total_Price: total, PayementID: payment, Status: "Cooking"}
	if err_order := conf.Query.Create(&order).Error; err_order != nil {
		return err_order
	}
	var last_order e.Order
	if err_last_order := conf.Query.Last(&last_order).Error; err_last_order != nil {
		return err_last_order
	}
	// fmt.Println(last_order.ID)
	for item := range cart {
		if err_order_detail := conf.Query.Create(&e.OrderDetail{OrderID: last_order.ID, MenuID: cart[item].ID_Menu, Quantity: cart[item].Quantity}).Error; err_order_detail != nil {
			return err_order_detail
		}
	}
	cart = nil
	return c.Redirect("/struk/" + strconv.Itoa(last_order.ID))
}

func Struk(c *fiber.Ctx) error {
	conf.Connect(".env")
	id_order := c.Params("id_order")
	var ordered Ordered
	var menuOrdered []MenuOrder
	//Query Get Menu Ordered
	conf.Query.Raw("SELECT tbl_menu.name as menu, tbl_menu.price * quantity as total_prices, quantity FROM tbl_order_detail JOIN tbl_menu ON tbl_order_detail.id_menu=tbl_menu.id WHERE id_order = ?", id_order).Scan(&menuOrdered)
	//Query Get Order
	conf.Query.Raw("SELECT tbl_meja.no_meja as no_meja, tbl_payement.name as payement, total_price, status FROM tbl_order JOIN tbl_meja ON tbl_meja.id=tbl_order.no_meja JOIN tbl_payement ON tbl_payement.id=tbl_order.id_payement WHERE tbl_order.id = ?", id_order).Scan(&ordered)

	return c.Render("struk", fiber.Map{
		"Title": "Struk",
		"Order": ordered,
		"Menu":  menuOrdered,
	})
}
