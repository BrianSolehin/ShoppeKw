package route

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	// Auth
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// Profile
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", controller.GetCurrentUser)
		auth.PUT("/profile", controller.UpdateCurrentUser)
		auth.DELETE("/profile", controller.DeleteCurrentUser)
	}

	// Admin
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RequireRole("admin"))
	{
		admin.GET("/users", controller.GetAllUsers)
		admin.DELETE("/users/:id", controller.DeleteUser)
		admin.GET("/sellers", controller.GetAllSellers) // ✅ Tambah ini
		admin.GET("/products", controller.GetAllProductsAdmin)
		admin.DELETE("/products/:id", controller.AdminDeleteProduct)
		admin.GET("/orders", controller.GetAllOrders)
		admin.GET("stores", controller.GetAllStores)       // Reuse dari buyer
		admin.DELETE("/stores/:id", controller.AdminDeleteStore)
		admin.GET("/transactions", controller.GetAllTransactions)

	}
}