package model

import "time"

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string
	Email    string    `gorm:"unique"`
	Password string
	Role     string    `gorm:"type:enum('admin','seller','buyer');default:'buyer'"`

	// Relasi
	Store     *Store     `gorm:"foreignKey:UserID"`       // Seller: satu user punya satu toko
	Orders    []Order    `gorm:"foreignKey:UserID"`       // Buyer: satu user bisa banyak order
	Products  []Product  `gorm:"foreignKey:UserID"`       // Seller: produk yang dijual

	CreatedAt time.Time
	UpdatedAt time.Time
}
