package repository

import (
	"ecommerce-api/config"
	"ecommerce-api/model"
)

func FindProductByID(id uint) (model.Product, error) {
	var prod model.Product
	err := config.DB.First(&prod, id).Error
	return prod, err
}

func CreateProduct(product *model.Product) error {
	return config.DB.Create(product).Error
}
