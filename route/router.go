package route

import (
	// "ecommerce-api/controller"
	// "ecommerce-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
			
	UserRoute(r)
	ProductRoute(r)
	SellerRoute(r)
	OrderRoute(r)
	StoreRoute(r) 
	TransactionRoute(r)

	return r
}

