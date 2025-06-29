package service

import (
	"errors"
	"ecommerce-api/dto/transaction"
	"ecommerce-api/model"
	"ecommerce-api/repository"
	"time"
)

type TransactionService interface {
	CreateTransaction(userID uint, req transaction.CreateRequest) (*model.Transaction, error)
	GetMyTransactions(userID uint) ([]model.Transaction, error)
	GetTransactionByID(id uint) (*model.Transaction, error)
	GetAllTransactions() ([]model.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	orderRepo       repository.OrderRepository
}

func NewTransactionService(
	tRepo repository.TransactionRepository,
	oRepo repository.OrderRepository,
) TransactionService {
	return &transactionService{
		transactionRepo: tRepo,
		orderRepo:       oRepo,
	}
}

func (s *transactionService) CreateTransaction(userID uint, req transaction.CreateRequest) (*model.Transaction, error) {
	order, err := s.orderRepo.GetByID(req.OrderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	if order.UserID != userID {
		return nil, errors.New("unauthorized: order does not belong to user")
	}

	if order.Status == "completed" {
		return nil, errors.New("order already completed")
	}

	transaction := &model.Transaction{
		OrderID:         req.OrderID,
		PaymentMethodID: req.PaymentMethodID,
		Status:          "pending", // manual langsung success
		TransactionDate: time.Now(),
	}

	err = s.transactionRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	order.Status = "completed"
	err = s.orderRepo.Update(order)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) GetMyTransactions(userID uint) ([]model.Transaction, error) {
	return s.transactionRepo.GetByUserID(userID)
}

func (s *transactionService) GetTransactionByID(id uint) (*model.Transaction, error) {
	return s.transactionRepo.GetByID(id)
}

func (s *transactionService) GetAllTransactions() ([]model.Transaction, error) {
	return s.transactionRepo.GetAll()
}
