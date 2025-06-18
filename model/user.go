package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string
	Email     string    `gorm:"unique"`
	Password  string
	Role      string    `gorm:"type:enum('admin','customer');default:'customer'"`
	Orders    []Order   `gorm:"foreignKey:UserID"`
	Products  []Product `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}