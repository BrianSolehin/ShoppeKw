package route

import (
	"ecommerce-api/controller"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(r *gin.Engine) {
	r.POST("/transactions", controller.CreateTransaction)
	r.GET("/transactions", controller.GetAllTransactions)
	r.GET("/transactions/:id", controller.GetTransactionByID)
}
