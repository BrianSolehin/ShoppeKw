package repository

import (
	"ecommerce-api/model"
	"ecommerce-api/config"
)

type OrderRepository interface {
	GetByID(id uint) (*model.Order, error)
	Update(order *model.Order) error
}


func FindAllOrders() ([]model.Order, error) {
	var orders []model.Order
	err := config.DB.Preload("User").Find(&orders).Error
	return orders, err
}

func FindOrderByID(id uint) (model.Order, error) {
	var order model.Order
	err := config.DB.Preload("User").First(&order, id).Error
	return order, err
}

func FindOrderDetailsByOrderID(orderID uint) ([]model.OrderDetail, error) {
	var details []model.OrderDetail
	err := config.DB.Where("order_id = ?", orderID).Preload("Product").Find(&details).Error
	return details, err
}

func DeleteOrderByID(id uint) error {
	return config.DB.Delete(&model.Order{}, id).Error
}

func CreateOrderWithDetails(order model.Order, details []model.OrderDetail) (uint, error) {
	tx := config.DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for i := range details {
		details[i].OrderID = order.ID
		if err := tx.Create(&details[i]).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()
	return order.ID, nil
}

func FindOrdersByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := config.DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func SaveProduct(product *model.Product) error {
	return config.DB.Save(product).Error
}

func DeleteProductByID(id uint) error {
	return config.DB.Delete(&model.Product{}, id).Error
}

func FindOrdersBySellerID(sellerID uint) ([]model.Order, error) {
	var orders []model.Order
	err := config.DB.
		Joins("JOIN order_details ON order_details.order_id = orders.id").
		Joins("JOIN products ON products.id = order_details.product_id").
		Joins("JOIN stores ON stores.id = products.store_id").
		Where("stores.user_id = ?", sellerID).
		Preload("User").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("OrderItems.Product.Category").
		Preload("OrderItems.Product.Store").
		Preload("OrderItems.Product.Store.User").
		Preload("Transaction").
		Find(&orders).Error

	return orders, err
}

func CheckSellerOwnsOrder(orderID uint, sellerID uint) (bool, error) {
	var count int64
	err := config.DB.
		Table("orders").
		Joins("JOIN order_details ON order_details.order_id = orders.id").
		Joins("JOIN products ON products.id = order_details.product_id").
		Joins("JOIN stores ON stores.id = products.store_id").
		Where("orders.id = ? AND stores.user_id = ?", orderID, sellerID).
		Count(&count).Error
	return count > 0, err
}

func UpdateOrderStatus(orderID uint, status string) error {
	return config.DB.Model(&model.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
