package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type FoodMaterialService interface {
	Create(request request.CreateFoodMaterial) (*model.FoodMaterial, error)
	GetAll() ([]model.FoodMaterial, error)
	GetByID(id uint) (*model.FoodMaterial, error)
	Update(id uint, request request.UpdateFoodMaterial) (*model.FoodMaterial, error)
	Delete(id uint) error
}

type foodMaterialService struct {
	foodMaterialRepo repo.FoodMaterialRepo
	validate         *validator.Validate
}

func NewFoodMaterialService(foodRepo repo.FoodMaterialRepo, validate *validator.Validate) FoodMaterialService {
	return &foodMaterialService{foodMaterialRepo: foodRepo, validate: validate}
}

func (s *foodMaterialService) Create(request request.CreateFoodMaterial) (*model.FoodMaterial, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newFoodMaterial := &model.FoodMaterial{
		Name:            request.Name,
		Unit:            request.Unit,
		StandardPerMeal: request.StandardPerMeal,
	}
	return s.foodMaterialRepo.Create(newFoodMaterial)
}

func (s *foodMaterialService) GetAll() ([]model.FoodMaterial, error) {
	return s.foodMaterialRepo.FindAll()
}

func (s *foodMaterialService) GetByID(id uint) (*model.FoodMaterial, error) {
	return s.foodMaterialRepo.FindByID(id)
}

func (s *foodMaterialService) Update(id uint, request request.UpdateFoodMaterial) (*model.FoodMaterial, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	foodMaterial, err := s.foodMaterialRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("food material not found")
	}

	foodMaterial.Name = request.Name
	foodMaterial.Unit = request.Unit
	foodMaterial.StandardPerMeal = request.StandardPerMeal
	return s.foodMaterialRepo.Update(foodMaterial)
}

func (s *foodMaterialService) Delete(id uint) error {
	return s.foodMaterialRepo.Delete(id)
}
