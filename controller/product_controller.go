package controller

import (
	"ecommerce-api/config"
	"ecommerce-api/dto/product"
	"ecommerce-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var req product.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		UserID:      req.UserID,
	}

	if err := config.DB.Create(&newProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil ditambahkan", "product": newProduct})
}

func GetAllProducts(c *gin.Context) {
	var products []model.Product
	config.DB.Preload("Category").Preload("User").Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product model.Product
	if err := config.DB.Preload("Category").Preload("User").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req product.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var prod model.Product
	if err := config.DB.First(&prod, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

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

	config.DB.Save(&prod)
	c.JSON(http.StatusOK, prod)
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&model.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}