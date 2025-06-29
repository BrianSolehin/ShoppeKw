package repository

import (
	"ecommerce-api/config"
	"ecommerce-api/model"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetByID(id uint) (*model.Transaction, error)
	GetByUserID(userID uint) ([]model.Transaction, error)
	GetAll() ([]model.Transaction, error)
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) Create(transaction *model.Transaction) error {
	return config.DB.Create(transaction).Error
}

func (r *transactionRepository) GetByID(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	err := config.DB.First(&transaction, id).Error
	return &transaction, err
}

func (r *transactionRepository) GetByUserID(userID uint) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := config.DB.
		Joins("JOIN orders ON orders.id = transactions.order_id").
		Where("orders.user_id = ?", userID).
		Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := config.DB.Find(&transactions).Error
	return transactions, err
}

type orderRepository struct{}

func NewOrderRepository() *orderRepository {
	return &orderRepository{}
}

func (r *orderRepository) GetByID(id uint) (*model.Order, error) {
	var order model.Order
	err := config.DB.First(&order, id).Error
	return &order, err
}

func (r *orderRepository) Update(order *model.Order) error {
	return config.DB.Save(order).Error
}
