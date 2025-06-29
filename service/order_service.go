package service

import (
	"errors"
	"strconv"
	"time" 
	"ecommerce-api/config"
	// "fmt"
	"gorm.io/gorm"
	"ecommerce-api/dto/order"  
	"ecommerce-api/model"      
	"ecommerce-api/repository" 
)

func GetAllOrders() ([]model.Order, error) {
	return repository.FindAllOrders()
}

func GetOrderByID(idStr string, userID uint, role string) (model.Order, []model.OrderDetail, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return model.Order{}, nil, errors.New("ID tidak valid")
	}

	order, err := repository.FindOrderByID(uint(id))
	if err != nil {
		return model.Order{}, nil, err
	}

	// ⛔ Cek kalau buyer, harus cocok userID-nya
	if role == "buyer" && order.UserID != userID {
		return model.Order{}, nil, errors.New("akses ditolak: bukan order milikmu")
	}

	details, err := repository.FindOrderDetailsByOrderID(uint(id))
	if err != nil {
		return model.Order{}, nil, err
	}

	return order, details, nil
}

func DeleteOrder(idStr string, userID uint, role string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("ID tidak valid")
	}

	order, err := repository.FindOrderByID(uint(id))
	if err != nil {
		return err
	}

	// Cek akses: buyer hanya bisa hapus order miliknya
	if role == "buyer" && order.UserID != userID {
		return errors.New("akses ditolak: kamu tidak bisa hapus order orang lain")
	}

	return repository.DeleteOrderByID(uint(id))
}

func CreateOrder(req order.CreateRequest, userID uint) (uint, error) {
	tx := config.DB.Begin()

	// Hitung total harga order
	var totalAmount float64
	for _, item := range req.Items {
		var product model.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			return 0, errors.New("produk tidak ditemukan")
		}

		if item.Quantity > product.Stock {
			tx.Rollback()
			return 0, errors.New("stok tidak cukup untuk produk ID: " + string(rune(item.ProductID)))
		}

		totalAmount += float64(item.Quantity) * product.Price
	}

	// Buat order utama
	order := model.Order{
		UserID:      userID,
		OrderDate:   time.Now(),
		TotalAmount: totalAmount,
		Status:      "completed", // bisa juga "pending"
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	// Simpan detail item dalam order
	for _, item := range req.Items {
		orderDetail := model.OrderDetail{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}

		if err := tx.Create(&orderDetail).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		// Kurangi stok produk
		if err := tx.Model(&model.Product{}).
			Where("id = ?", item.ProductID).
			Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()
	return order.ID, nil
}

func GetMyOrders(userID uint) ([]model.Order, error) {
	return repository.FindOrdersByUserID(userID)
}
func GetOrdersBySeller(userID uint) ([]model.Order, error) {
	return repository.FindOrdersBySellerID(userID)
}

func UpdateOrderStatusBySeller(orderID uint, sellerID uint, newStatus string) error {
	// validasi status
	if newStatus != "completed" && newStatus != "cancelled" && newStatus != "pending" {
		return errors.New("status tidak valid")
	}

	// cek apakah seller punya produk di order tsb
	hasAccess, err := repository.CheckSellerOwnsOrder(orderID, sellerID)
	if err != nil {
		return err
	}
	if !hasAccess {
		return errors.New("akses ditolak: order bukan milikmu")
	}

	return repository.UpdateOrderStatus(orderID, newStatus)
}
