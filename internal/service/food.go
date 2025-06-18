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

func (s *foodService) Create(request request.CreateMealItem) (*model.food, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newItem := &model.Food{
		MealTypeID: request.MealTypeID,
		FoodID:     request.FoodID,
		Quantity:   request.Quantity,
	}

	return s.foodRepo.Create(newItem)
}

func (s *foodService) GetAll() ([]model.Food, error) {
	return s.mealItemRepo.FindAll()
}

func (s *foodService) GetByID(id uint) (*model.Food, error) {
	return s.foodRepo.FindByID(id)
}

func (s *foodService) Update(id uint, request request.UpdateMealItem) (*model.Food, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	item, err := s.foodRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	item.MealTypeID = request.MealTypeID
	item.FoodID = request.FoodID
	item.Quantity = request.Quantity
	item.MealType = nil
	item.Food = nil
	return s.foodRepo.Update(item)
}

func (s *foodService) Delete(id uint) error {
	return s.foodRepo.Delete(id)
}
