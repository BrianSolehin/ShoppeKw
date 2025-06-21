package controller

import "github.com/gin-gonic/gin"

// GET /admin/stores
func GetAllStores(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Semua store ditampilkan (admin only)"})
}

// GET /seller/my-store
func GetMyStore(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Ini store saya (seller only)"})
}
