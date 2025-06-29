package controller

import (
	"ecommerce-api/config"
	"ecommerce-api/model"
	"ecommerce-api/repository"
	"ecommerce-api/service"
	userdto "ecommerce-api/dto/user"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register user baru
func Register(c *gin.Context) {
	var req userdto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.NewUserService(repository.NewUserRepository(config.DB))
	resp, err := userService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil", "user": resp})
}
// Login user
func Login(c *gin.Context) {
	var req userdto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.NewUserService(repository.NewUserRepository(config.DB))
	token, err := userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   token,
	})
}
// Ambil data user yang sedang login
func GetCurrentUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}

// Update user yang sedang login
func UpdateCurrentUser(c *gin.Context) {
	currentUser := c.MustGet("user").(model.User)

	var req userdto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name != "" {
		currentUser.Name = req.Name
	}
	if req.Email != "" {
		currentUser.Email = req.Email
	}
	if req.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		currentUser.Password = string(hashed)
	}

	if err := config.DB.Save(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	currentUser.Password = ""
	c.JSON(http.StatusOK, currentUser)
}

// Hapus user yang sedang login
func DeleteCurrentUser(c *gin.Context) {
	currentUser := c.MustGet("user").(model.User)

	if err := config.DB.Delete(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Akun berhasil dihapus"})
}

func GetAllUsers(c *gin.Context) {
	users, err := service.NewUserService(repository.NewUserRepository(config.DB)).GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.NewUserService(repository.NewUserRepository(config.DB)).DeleteUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}

func GetAllSellers(c *gin.Context) {
	var sellers []model.User
	if err := config.DB.Where("role = ?", "seller").Find(&sellers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hapus password biar gak ikut ke-expose
	for i := range sellers {
		sellers[i].Password = ""
	}

	c.JSON(http.StatusOK, sellers)
}