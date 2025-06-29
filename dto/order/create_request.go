package order

type OrderItem struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	// Price     float64 `json:"price" binding:"required"` // ✅ Tambahkan baris ini
}

type CreateRequest struct {
	// UserID uint        `json:"user_id" binding:"required"`
	Items  []OrderItem `json:"items" binding:"required,dive,required"`
}