package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(request request.Register) (*model.User, error)
	SignIn(request request.SignIn) (*model.User, error)
}

type authService struct {
	userRepo repo.UserRepo
	validate *validator.Validate
}

func NewAuthService(userRepo repo.UserRepo, validate *validator.Validate) AuthService {
	return &authService{userRepo: userRepo, validate: validate}
}

func (s *authService) Register(request request.Register) (*model.User, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return nil, err
	}

	existingUser, _ := s.userRepo.FindByUsername(request.Username)
	if existingUser != nil {
		return nil, errors.New("user with this username already exists")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	newUser := &model.User{
		Username:     request.Username,
		Name:         request.Name,
		PasswordHash: string(hashedPassword),
		Role:         request.Role,
	}

	newUser, err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *authService) SignIn(request request.SignIn) (*model.User, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByUsername(request.Username)
	if err != nil {
		return nil, errors.New("Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, errors.New("Invalid username or password")
		}
		return nil, err
	}
	return user, nil
}
