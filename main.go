package main

import (
	"ecommerce-api/config"
	"ecommerce-api/model"
	"ecommerce-api/route"
)

func main() {
	// ✅ KONEKSI DATABASE
	config.ConnectDB()

	// ✅ MIGRASI DATABASE
	config.DB.AutoMigrate(
		&model.User{},
		&model.Store{},
		&model.Category{},
		&model.Product{},
		&model.Order{},
		&model.OrderDetail{},
		&model.PaymentMethod{},
		&model.Transaction{},
	)

	// ✅ INISIALISASI ROUTER LENGKAP DENGAN MIDDLEWARE & PREFIX
	r := route.SetupRouter()

	// ✅ JALANKAN SERVER
	r.Run("0.0.0.0:8080")

}
