package config

import (
	"fmt"
	"os"
	"time"
	"github.com/joho/godotenv" 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

)

var DB *gorm.DB

func ConnectDB() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load file .env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("REAL DSN YANG DIKIRIM =", dsn) // ⬅️
	fmt.Println("DB_HOST =", os.Getenv("DB_HOST"))  // ⬅️
	for i := 1; i <= 1; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ Berhasil konek ke database")
			return
		}
		fmt.Printf("Percobaan %d: ❌ Gagal konek database: %s\n", i, err)
		time.Sleep(2 * time.Second)
	}
	panic("❌ Gagal konek ke database setelah 10 percobaan")
}
