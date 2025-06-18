package model

import "time"

type Order struct {
	ID          uint          `gorm:"primaryKey"`
	UserID      uint
	User        User
	OrderDate   time.Time
	Status      string         `gorm:"type:enum('pending','completed','cancelled');default:'pending'"`
	TotalAmount float64
	OrderItems  []OrderDetail  `gorm:"foreignKey:OrderID"`
	Transaction Transaction
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
