package service

import (
	"ecommerce-api/model"
	"ecommerce-api/repository"
	"errors"
	"ecommerce-api/config"
)

func GetMyStore(userID uint) (model.Store, error) {
	return repository.FindStoreByUserID(userID)
}

func CreateStore(name string, userID uint) (model.Store, error) {
	store := model.Store{
		Name:   name,
		UserID: userID,
	}

	if err := config.DB.Create(&store).Error; err != nil {
		return model.Store{}, err
	}

	// 🔥 Ambil ulang dengan preload biar relasi keisi
	var result model.Store
	err := config.DB.
		Preload("User").
		Preload("Products").
		First(&result, store.ID).Error

	return result, err
}


func GetAllStores() ([]model.Store, error) {
	return repository.FindAllStores()
}



func DeleteStoreByID(storeID uint, sellerID uint) error {
	store, err := repository.FindStoreByID(storeID)
	if err != nil {
		return err
	}

	if store.UserID != sellerID {
		return errors.New("akses ditolak: toko bukan milikmu")
	}

	return repository.DeleteStoreByID(storeID)
}

func UpdateStoreByID(storeID uint, sellerID uint, newName string) error {
    // Cek apakah store milik seller yang login
    store, err := repository.FindStoreByID(storeID)
    if err != nil {
        return err
    }

    if store.UserID != sellerID {
        return errors.New("akses ditolak: toko bukan milik Anda")
    }

    store.Name = newName
    return repository.UpdateStore(&store)
}

