package main

import (
	"github.com/gin-gonic/gin"
	"ecommerce-api/config"       // koneksi DB
	"ecommerce-api/model"  
	// "ecommerce-api/seed"  
	"ecommerce-api/route" // ⬅️ tambahkan baris ini
)

func main() {
	r := gin.Default()

	// ✅ KONEKSI DATABASE
	config.ConnectDB()
	route.TransactionRoute(r)
	route.OrderRoute(r)
	route.UserRoute(r)
	route.ProductRoute(r)

	// ✅ MIGRATE SEMUA TABEL
	config.DB.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Product{},
		&model.Order{},
		&model.OrderDetail{},
		&model.PaymentMethod{},
		&model.Transaction{},
	)

	// // seed.InsertDummyData()

	// Endpoint tes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is working with DB!",
		})
	})

	// Jalankan server di localhost:8080
	r.Run(":8080")
}
