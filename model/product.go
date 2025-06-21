package model

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryID  uint
	UserID      uint      // Masih boleh kalau kamu mau track siapa yang buat produk
	StoreID     uint      // Tambahkan ini untuk relasi ke Store
	Store       Store     `gorm:"foreignKey:StoreID"` // Tambahkan ini untuk eager loading

	CreatedAt   time.Time
	UpdatedAt   time.Time
}
