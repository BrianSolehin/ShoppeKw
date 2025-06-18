package model

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryID  uint
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
