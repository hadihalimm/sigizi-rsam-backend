package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type FoodService interface {
	Create(request request.CreateFood) (*model.Food, error)
	GetAll() ([]model.Food, error)
	GetByID(id uint) (*model.Food, error)
	Update(id uint, request request.UpdateFood) (*model.Food, error)
	Delete(id uint) error
}

type foodService struct {
	foodRepo repo.FoodRepo
	validate *validator.Validate
}

func NewFoodService(foodRepo repo.FoodRepo,
	validate *validator.Validate) FoodService {
	return &foodService{foodRepo: foodRepo, validate: validate}
}

func (s *foodService) Create(request request.CreateFood) (*model.Food, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	var usages []model.FoodMaterialUsage
	for _, usage := range request.FoodMaterialUsages {
		usages = append(usages, model.FoodMaterialUsage{
			FoodMaterialID: usage.FoodMaterialID,
			QuantityUsed:   usage.QuantityUsed,
		})
	}

	newFood := &model.Food{
		Name:               request.Name,
		FoodMaterialUsages: usages,
	}

	return s.foodRepo.Create(newFood)
}

func (s *foodService) GetAll() ([]model.Food, error) {
	return s.foodRepo.FindAll()
}

func (s *foodService) GetByID(id uint) (*model.Food, error) {
	return s.foodRepo.FindByID(id)
}

func (s *foodService) Update(id uint, request request.UpdateFood) (*model.Food, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	food, err := s.foodRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	food.Name = request.Name
	food.FoodMaterialUsages = []model.FoodMaterialUsage{}
	var usages []model.FoodMaterialUsage
	for _, usage := range request.FoodMaterialUsages {
		usages = append(usages, model.FoodMaterialUsage{
			FoodID:         id, // optional — can also be set in repo
			FoodMaterialID: usage.FoodMaterialID,
			QuantityUsed:   usage.QuantityUsed,
		})
	}
	food.FoodMaterialUsages = usages

	return s.foodRepo.Update(food)
}

func (s *foodService) Delete(id uint) error {
	return s.foodRepo.Delete(id)
}
