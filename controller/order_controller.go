// controller/order_controller.go
package controller

import (
	"ecommerce-api/config"
	"ecommerce-api/dto/order"
	"ecommerce-api/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var req order.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mulai transaksi
	tx := config.DB.Begin()

	total := 0.0
	order := model.Order{
		UserID:      req.UserID,
		Status:      "completed",
		OrderDate:   time.Now(),
		TotalAmount: 0,
	}

	// Simpan order dulu
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop items & simpan order_detail
	for _, item := range req.Items {
		var product model.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		orderDetail := model.OrderDetail{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}

		if err := tx.Create(&orderDetail).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		total += product.Price * float64(item.Quantity)
	}

	// Update total harga
	order.TotalAmount = total
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit jika semua aman
	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"message": "Order berhasil dibuat", "order_id": order.ID})
}

func GetAllOrders(c *gin.Context) {
	var orders []model.Order
	config.DB.Preload("User").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var order model.Order
	if err := config.DB.Preload("User").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order tidak ditemukan"})
		return
	}
	var details []model.OrderDetail
	config.DB.Where("order_id = ?", order.ID).Preload("Product").Find(&details)

	c.JSON(http.StatusOK, gin.H{
		"order":  order,
		"detail": details,
	})
}

func DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&model.Order{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order berhasil dihapus"})
}