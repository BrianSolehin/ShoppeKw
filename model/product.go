package model

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryID  uint
	Category    Category  `gorm:"foreignKey:CategoryID"` // ✅ Tambahkan ini
	UserID      uint
	StoreID     uint
	Store       Store     `gorm:"foreignKey:StoreID"`

	CreatedAt   time.Time
	UpdatedAt   time.Time
}
