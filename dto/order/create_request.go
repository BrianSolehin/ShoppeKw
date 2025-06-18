package order

type OrderItem struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type CreateRequest struct {
	UserID uint        `json:"user_id" binding:"required"`
	Items  []OrderItem `json:"items" binding:"required,dive,required"`
}