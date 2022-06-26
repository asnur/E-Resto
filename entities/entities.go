package entities

//Table Meja
type Tabler interface {
	TableName() string
}
type Table struct {
	ID      int `gorm:"primary_key"`
	No_Meja int
}

func (t *Table) TableName() string {
	return "tbl_meja"
}

//Table Order
type StatusOrders string

const (
	Ready   StatusOrders = "Ready"
	Cooking StatusOrders = "Cooking"
	Done    StatusOrders = "Done"
)

type Order struct {
	ID           int   `gorm:"primary_key"`
	TableNo_Meja int   `gorm:"column:no_meja"`
	Table        Table `gorm:"foreignkey:TableNo_Meja"`
	Total_Price  int
	PayementID   int           `gorm:"column:id_payement"`
	Payement     Payement      `gorm:"foreignkey:PayementID"`
	Status       StatusOrders  `sql:"type:status_orders"`
	OrderDetails []OrderDetail `gorm:"foreignkey:OrderID"`
}

func (t *Order) TableName() string {
	return "tbl_order"
}

//Table Order Detail
type OrderDetail struct {
	ID       int   `gorm:"primary_key"`
	OrderID  int   `gorm:"column:id_order"`
	Order    Order `gorm:"foreignkey:OrderID"`
	Quantity int
}

func (t *OrderDetail) TableName() string {
	return "tbl_order_detail"
}

//Table Menu
type Menu struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Price      int
	CategoryID int      `gorm:"column:id_category"`
	Category   Category `gorm:"foreignkey:CategoryID"`
	Foto       string
}

func (t *Menu) TableName() string {
	return "tbl_menu"
}

//Table Category
type Category struct {
	ID   int `gorm:"primary_key"`
	Name string
	Icon string
}

func (t *Category) TableName() string {
	return "tbl_category"
}

//Table Payement
type Payement struct {
	ID   int `gorm:"primary_key"`
	Name string
}

func (t *Payement) TableName() string {
	return "tbl_payement"
}
