package route

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoute(r *gin.Engine) {
	order := r.Group("/orders")
	order.Use(middleware.AuthMiddleware())
	{
		order.POST("", controller.CreateOrder)           // Buyer only
		order.GET("/:id", controller.GetOrderByID)       // Buyer only
		order.DELETE("/:id", controller.DeleteOrder)     // Buyer only
	}

	// GET semua order: hanya admin
	
}
