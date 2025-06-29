package main

import (
	"ecommerce-api/config"
	"ecommerce-api/model"
	"ecommerce-api/route"
	"log"
	"fmt"
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
	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
	fmt.Println("User routes registered...")
	// ✅ JALANKAN SERVER
	r.Run(":8080")

}
