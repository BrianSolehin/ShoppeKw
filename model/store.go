package model

import "time"

type Store struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`               // Foreign key ke users.id
	User      User      `gorm:"foreignKey:UserID"`      // Relasi balik ke user (pemilik toko)
	Products  []Product `gorm:"foreignKey:StoreID"`     // Semua produk milik toko ini

	CreatedAt time.Time
	UpdatedAt time.Time
}
