package controller

import (
	"ecommerce-api/dto/transaction"
	"ecommerce-api/repository"
	"ecommerce-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var txRepo = repository.NewTransactionRepository()
var orderRepo = repository.NewOrderRepository()
var txService = service.NewTransactionService(txRepo, orderRepo)

func CreateTransaction(c *gin.Context) {
	var req transaction.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	tx, err := txService.CreateTransaction(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

func GetMyTransactions(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)
	txs, err := txService.GetMyTransactions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, txs)
}

func GetTransactionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tx, err := txService.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	c.JSON(http.StatusOK, tx)
}

func GetAllTransactions(c *gin.Context) {
	txs, err := txService.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, txs)
}
