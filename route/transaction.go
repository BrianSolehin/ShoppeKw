package route

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(r *gin.Engine) {
	// Buyer transaction routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(), middleware.RequireRole("buyer"))
	{
		auth.POST("/transactions", controller.CreateTransaction)
		auth.GET("/transactions", controller.GetMyTransactions)
		auth.GET("/transactions/:id", controller.GetTransactionByID)
	}
}
