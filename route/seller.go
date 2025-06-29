package route


import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gin-gonic/gin"
)



func SellerRoute(r *gin.Engine) {
	seller := r.Group("/seller")
	seller.Use(middleware.AuthMiddleware(), middleware.RequireRole("seller"))
	{
		seller.GET("/profile", controller.GetCurrentUser)
		seller.PUT("/profile", controller.UpdateCurrentUser)
		seller.POST("/products", controller.CreateProduct)
		seller.PUT("/products/:id", controller.UpdateProduct)
		seller.DELETE("/products/:id", controller.DeleteProduct)
		seller.GET("/orders", controller.GetOrdersBySeller)
    	seller.PUT("/orders/:id/status", controller.UpdateOrderStatusBySeller)
		seller.GET("/stores", controller.GetMyStore)
		seller.POST("/stores", controller.CreateStore)
		seller.PUT("/stores/:id", controller.UpdateStoreByID)   // update store by ID
		seller.DELETE("/stores/:id", controller.DeleteStoreByID) // delete store by ID
	}
}
