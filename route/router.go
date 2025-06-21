package route

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// Public: bisa dilihat siapa aja
	r.GET("/products", controller.GetAllProducts)
	r.GET("/products/:id", controller.GetProductByID)

	// Admin-only routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RequireRole("admin"))
	{
		admin.GET("/users", controller.GetAllUsers)
		admin.DELETE("/users/:id", controller.DeleteUser)
		admin.GET("/stores", controller.GetAllStores)
		admin.GET("/orders", controller.GetAllOrders)
	}

	// Seller-only routes
	seller := r.Group("/seller")
	seller.Use(middleware.AuthMiddleware(), middleware.RequireRole("seller"))
	{
		seller.POST("/products", controller.CreateProduct)
		seller.PUT("/products/:id", controller.UpdateProduct)
		seller.DELETE("/products/:id", controller.DeleteProduct)
		seller.GET("/my-store", controller.GetMyStore)
	}

	// Buyer-only routes
	buyer := r.Group("/buyer")
	buyer.Use(middleware.AuthMiddleware(), middleware.RequireRole("buyer"))
	{
		buyer.POST("/orders", controller.CreateOrder)
		buyer.GET("/orders", controller.GetMyOrders)
		buyer.POST("/transactions", controller.CreateTransaction)
	}

	return r
}
