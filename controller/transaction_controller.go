package controller

import (
	"ecommerce-api/config"
	"ecommerce-api/dto/transaction"
	"ecommerce-api/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var req transaction.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTx := model.Transaction{
		OrderID:         req.OrderID,
		PaymentMethodID: req.PaymentMethodID,
		Status:          req.Status,
		TransactionDate: time.Now(),
	}

	if err := config.DB.Create(&newTx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaksi berhasil", "transaction_id": newTx.ID})
}

func GetAllTransactions(c *gin.Context) {
	var txs []model.Transaction
	config.DB.Preload("Order").Preload("PaymentMethod").Find(&txs)
	c.JSON(http.StatusOK, txs)
}

func GetTransactionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var tx model.Transaction
	if err := config.DB.Preload("Order").Preload("PaymentMethod").First(&tx, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, tx)
}
