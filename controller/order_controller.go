// controller/order_controller.go
package controller

import (
	"ecommerce-api/dto/order"
	"net/http"
	"github.com/gin-gonic/gin"
	"ecommerce-api/service"
)

func CreateOrder(c *gin.Context) {
	var req order.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID, err := service.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order berhasil dibuat", "order_id": orderID})
}

func GetAllOrders(c *gin.Context) {
	orders, err := service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	userIDInterface, _ := c.Get("userID")
	roleInterface, _ := c.Get("role")

	userID := userIDInterface.(uint)
	role := roleInterface.(string)

	order, detail, err := service.GetOrderByID(id, userID, role)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order, "detail": detail})
}


func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	userIDInterface, _ := c.Get("userID")
	roleInterface, _ := c.Get("role")

	userID := userIDInterface.(uint)
	role := roleInterface.(string)

	err := service.DeleteOrder(id, userID, role)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order berhasil dihapus"})
}


func GetMyOrders(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan dalam token"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID tidak valid"})
		return
	}

	orders, err := service.GetMyOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}