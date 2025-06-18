package model

import "time"

type Transaction struct {
	ID              uint      `gorm:"primaryKey"`
	OrderID         uint
	PaymentMethodID uint
	Status          string    `gorm:"type:enum('success','failed','pending');default:'pending'"`
	TransactionDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
