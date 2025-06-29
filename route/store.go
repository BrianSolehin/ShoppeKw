package route

import (
	"ecommerce-api/controller"
	"github.com/gin-gonic/gin"
)

func StoreRoute(r *gin.Engine) {
	// Untuk buyer & public
	r.GET("/stores", controller.GetAllStores)
	
}
