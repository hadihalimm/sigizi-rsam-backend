package service

import (
	"errors"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	Update(id uint, request request.UpdateUser) (*model.User, error)
	Delete(id uint) error
	ResetPassword(id uint) (*model.User, error)
	UpdatePassword(id uint, request request.UpdatePassword) (*model.User, error)
	UpdateName(id uint, request request.UpdateName) (*model.User, error)
}

type userService struct {
	userRepo repo.UserRepo
	validate *validator.Validate
}

func NewUserService(userRepo repo.UserRepo, validate *validator.Validate) UserService {
	return &userService{userRepo: userRepo, validate: validate}
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) Update(id uint, request request.UpdateUser) (*model.User, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	user.Username = request.Username
	user.Name = request.Name
	user.Role = request.Role

	return s.userRepo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *userService) ResetPassword(id uint) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("DEFAULT_PASSWORD")), bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)
	return s.userRepo.Update(user)
}

func (s *userService) UpdatePassword(id uint, request request.UpdatePassword) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)
	return s.userRepo.Update(user)
}

func (s *userService) UpdateName(id uint, request request.UpdateName) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	user.Name = request.Name
	return s.userRepo.Update(user)
}
