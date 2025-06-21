package repository

import (
	"ecommerce-api/model"
	"ecommerce-api/config"
)

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
