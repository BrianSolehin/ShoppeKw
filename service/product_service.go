package service

import (
	"errors"
	"ecommerce-api/config"
	"ecommerce-api/model"
	"ecommerce-api/dto/product"
	"ecommerce-api/repository"
)

func CreateProduct(req product.CreateRequest, userID uint) (model.Product, error) {
	var store model.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return model.Product{}, errors.New("store tidak ditemukan")
	}

	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		UserID:      userID,
		StoreID:     store.ID, // ✅ ini penting
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func UpdateProduct(id uint, req product.UpdateRequest, userID uint, role string) (model.Product, error) {
	prod, err := repository.FindProductByID(id)
	if err != nil {
		return prod, errors.New("produk tidak ditemukan")
	}

	// 🔐 Cek role seller
	if role == "seller" && prod.UserID != userID {
		return prod, errors.New("kamu tidak punya akses ke produk ini")
	}

	// 🛠 Update data
	if req.Name != "" {
		prod.Name = req.Name
	}
	if req.Description != "" {
		prod.Description = req.Description
	}
	if req.Price != 0 {
		prod.Price = req.Price
	}
	if req.Stock != 0 {
		prod.Stock = req.Stock
	}
	if req.CategoryID != 0 {
		prod.CategoryID = req.CategoryID
	}

	err = repository.SaveProduct(&prod)
	return prod, err
}

func DeleteProduct(id uint, userID uint, role string) error {
	prod, err := repository.FindProductByID(id)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}

	if role == "seller" && prod.UserID != userID {
		return errors.New("kamu tidak punya akses ke produk ini")
	}

	return repository.DeleteProductByID(id)
}

// service/product_service.go
func ForceDeleteProduct(id uint) error {
	_, err := repository.FindProductByID(id)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}
	return repository.DeleteProductByID(id)
}

func SaveProduct(product *model.Product) error {
	return config.DB.Save(product).Error
}

func DeleteProductByID(id uint) error {
	return config.DB.Delete(&model.Product{}, id).Error
}