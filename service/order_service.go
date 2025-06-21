package service

import (
	"errors"
	"strconv"
	"time" // ✅ UNTUK time.Now()

	"ecommerce-api/dto/order"  // ✅ UNTUK order.CreateRequest
	"ecommerce-api/model"      // ✅ UNTUK model.Order, model.OrderDetail
	"ecommerce-api/repository" // ✅ UNTUK repository.CreateOrderWithDetails
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


func CreateOrder(req order.CreateRequest) (uint, error) {
	if req.UserID == 0 || len(req.Items) == 0 {
		return 0, errors.New("data order tidak lengkap")
	}

	order := model.Order{
		UserID:      req.UserID,
		Status:      "completed",
		OrderDate:   time.Now(),
		TotalAmount: 0,
	}

	// Hitung total dan simpan order detail
	var total float64
	var details []model.OrderDetail
	for _, item := range req.Items {
		detail := model.OrderDetail{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		total += item.Price * float64(item.Quantity)
		details = append(details, detail)
	}
	order.TotalAmount = total

	orderID, err := repository.CreateOrderWithDetails(order, details)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func GetMyOrders(userID uint) ([]model.Order, error) {
	return repository.FindOrdersByUserID(userID)
}