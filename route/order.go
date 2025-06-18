package route

import (
	"ecommerce-api/controller"
	"github.com/gin-gonic/gin"
)

func OrderRoute(r *gin.Engine) {
	r.POST("/orders", controller.CreateOrder)
	r.GET("/orders", controller.GetAllOrders)
	r.GET("/orders/:id", controller.GetOrderByID)
	r.DELETE("/orders/:id", controller.DeleteOrder)
}
