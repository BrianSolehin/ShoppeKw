package transaction

type CreateRequest struct {
	OrderID         uint   `json:"order_id" binding:"required"`
	PaymentMethodID uint   `json:"payment_method_id" binding:"required"`
	Status          string `json:"status" binding:"required"`
}
