package repository

import (
	"ecommerce-api/config"
	"ecommerce-api/model"
)

func FindStoreByUserID(userID uint) (model.Store, error) {
	var store model.Store
	err := config.DB.
		Preload("User").
		Preload("Products").
		Preload("Products.Category").
		Preload("Products.Store").
		Preload("Products.Store.User").
		Where("user_id = ?", userID).
		First(&store).Error
	return store, err
}


func CreateStore(store model.Store) (model.Store, error) {
	err := config.DB.Create(&store).Error
	return store, err
}

func FindAllStores() ([]model.Store, error) {
	var stores []model.Store
	err := config.DB.Preload("User").Preload("Products").Find(&stores).Error
	return stores, err
}



func SaveStore(store *model.Store) error {
	return config.DB.Save(store).Error
}


func DeleteStoreByID(id uint) error {
	return config.DB.Delete(&model.Store{}, id).Error
}

func FindStoreByID(id uint) (model.Store, error) {
	var store model.Store
	err := config.DB.First(&store, id).Error
	return store, err
}

func UpdateStore(store *model.Store) error {
    return config.DB.Save(store).Error
}
