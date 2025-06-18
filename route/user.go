package route

import (
	"ecommerce-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	// Public routes
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// Semua route bisa diakses tanpa token
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:id", controller.GetUserByID)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
