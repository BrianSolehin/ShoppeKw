package model

import "time"

type PaymentMethod struct {
	ID           uint          `gorm:"primaryKey"`
	Name         string
	Transactions []Transaction `gorm:"foreignKey:PaymentMethodID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
