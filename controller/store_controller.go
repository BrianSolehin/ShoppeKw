package controller

import (
	"ecommerce-api/service"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"ecommerce-api/repository" 

)

func GetMyStore(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	store, err := service.GetMyStore(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, store)
}

func CreateStore(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store, err := service.CreateStore(req.Name, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, store)
}

func GetAllStores(c *gin.Context) {
	stores, err := service.GetAllStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stores)
}






func DeleteStoreByID(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID toko tidak valid"})
		return
	}

	sellerID := c.MustGet("userID").(uint)
	err = service.DeleteStoreByID(uint(storeID), sellerID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Toko berhasil dihapus"})
}


func UpdateStoreByID(c *gin.Context) {
    storeIDStr := c.Param("id")
    storeID, err := strconv.Atoi(storeIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID toko tidak valid"})
        return
    }

    var req struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    sellerID := c.MustGet("userID").(uint)

    err = service.UpdateStoreByID(uint(storeID), sellerID, req.Name)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Store berhasil diperbarui"})
}

func AdminDeleteStore(c *gin.Context) {
	idParam := c.Param("id")
	storeID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	err = repository.DeleteStoreByID(uint(storeID)) // Tanpa cek pemilik
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Toko berhasil dihapus oleh admin"})
}
