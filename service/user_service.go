package service

import (
	"ecommerce-api/dto/user"
	"ecommerce-api/model"
	"ecommerce-api/repository"
	"ecommerce-api/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Interface
type UserService interface {
	Register(req user.RegisterRequest) (*user.UserResponse, error)
	Login(req user.LoginRequest) (string, error)
	GetAllUsers() ([]model.User, error)
	DeleteUserByID(id uint) error
}

// Struct implementasi interface
type userService struct {
	repo repository.UserRepository
}

// Konstruktor
func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

// Register logic
func (s *userService) Register(req user.RegisterRequest) (*user.UserResponse, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	newUser := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	err := s.repo.Create(&newUser)
	if err != nil {
		return nil, err
	}

	return &user.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
		Role:  newUser.Role,
	}, nil
}

// Login logic
func (s *userService) Login(req user.LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("email tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("password salah")
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", errors.New("gagal membuat token")
	}

	return token, nil
}

// Admin: Get all users
func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}

// Admin: Delete user by ID
func (s *userService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
