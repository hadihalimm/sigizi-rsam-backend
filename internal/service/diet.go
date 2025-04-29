package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type DietService interface {
	Create(request request.CreateDiet) (*model.Diet, error)
	GetAll() ([]model.Diet, error)
	GetByID(id uint) (*model.Diet, error)
	Update(id uint, request request.UpdateDiet) (*model.Diet, error)
	Delete(id uint) error
}

type dietService struct {
	dietRepo repo.DietRepo
	validate *validator.Validate
}

func NewDietService(dietRepo repo.DietRepo, validate *validator.Validate) DietService {
	return &dietService{dietRepo: dietRepo, validate: validate}
}

func (s *dietService) Create(request request.CreateDiet) (*model.Diet, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newDiet := &model.Diet{
		Code: request.Code,
		Name: request.Name,
	}
	return s.dietRepo.Create(newDiet)
}

func (s *dietService) GetAll() ([]model.Diet, error) {
	return s.dietRepo.FindAll()
}

func (s *dietService) GetByID(id uint) (*model.Diet, error) {
	return s.dietRepo.FindByID(id)
}

func (s *dietService) Update(id uint, request request.UpdateDiet) (*model.Diet, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	diet, err := s.dietRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("diet not found")
	}

	diet.Code = request.Code
	diet.Name = request.Name
	return s.dietRepo.Update(diet)
}

func (s *dietService) Delete(id uint) error {
	return s.dietRepo.Delete(id)
}
