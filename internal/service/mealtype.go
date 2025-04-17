package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type MealTypeService interface {
	Create(request request.CreateMealType) (*model.MealType, error)
	GetAll() ([]model.MealType, error)
	GetByID(id uint) (*model.MealType, error)
	Update(id uint, request request.UpdateMealType) (*model.MealType, error)
	Delete(id uint) error
}

type mealTypeService struct {
	mealTypeRepo repo.MealTypeRepo
	validate     *validator.Validate
}

func NewMealTypeService(mealTypeRepo repo.MealTypeRepo, validate *validator.Validate) MealTypeService {
	return &mealTypeService{mealTypeRepo: mealTypeRepo, validate: validate}
}

func (s *mealTypeService) Create(request request.CreateMealType) (*model.MealType, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newMealType := &model.MealType{
		Name: request.Name,
	}
	return s.mealTypeRepo.Create(newMealType)
}

func (s *mealTypeService) GetAll() ([]model.MealType, error) {
	return s.mealTypeRepo.FindAll()
}

func (s *mealTypeService) GetByID(id uint) (*model.MealType, error) {
	return s.mealTypeRepo.FindByID(id)
}

func (s *mealTypeService) Update(id uint, request request.UpdateMealType) (*model.MealType, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	mt, err := s.mealTypeRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("food not found")
	}

	mt.Name = request.Name
	return s.mealTypeRepo.Update(mt)
}

func (s *mealTypeService) Delete(id uint) error {
	return s.mealTypeRepo.Delete(id)
}
