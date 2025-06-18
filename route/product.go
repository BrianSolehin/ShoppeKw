package route

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoute(r *gin.Engine) {
	// 🟢 PUBLIC: Lihat produk tanpa login
	r.GET("/products", controller.GetAllProducts)
	r.GET("/products/:id", controller.GetProductByID)

	// 🔐 PROTECTED: Butuh login + role admin
	auth := r.Group("/products")
	auth.Use(middleware.AuthMiddleware()) // cek token
	{
		auth.POST("", middleware.IsAdmin(), controller.CreateProduct)
		auth.PUT("/:id", middleware.IsAdmin(), controller.UpdateProduct)
		auth.DELETE("/:id", middleware.IsAdmin(), controller.DeleteProduct)
	}
}
