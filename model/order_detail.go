package model

type OrderDetail struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Product   Product
	Quantity  int
	Price     float64
}
	