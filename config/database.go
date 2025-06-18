package config

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // DB adalah variable global untuk menyimpan koneksi database

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	fmt.Println("✅ Database terkoneksi")
	DB = database
}