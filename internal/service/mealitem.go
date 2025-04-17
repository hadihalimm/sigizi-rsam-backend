package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type MealItemService interface {
	Create(request request.CreateMealItem) (*model.MealItem, error)
	GetAll() ([]model.MealItem, error)
	GetByID(id uint) (*model.MealItem, error)
	Update(id uint, request request.UpdateMealItem) (*model.MealItem, error)
	Delete(id uint) error
}

type mealItemService struct {
	mealItemRepo repo.MealItemRepo
	validate     *validator.Validate
}

func NewMealItemService(mealItemRepo repo.MealItemRepo,
	validate *validator.Validate) MealItemService {
	return &mealItemService{mealItemRepo: mealItemRepo, validate: validate}
}

func (s *mealItemService) Create(request request.CreateMealItem) (*model.MealItem, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newItem := &model.MealItem{
		Date:       request.Date,
		MealTypeID: request.MealTypeID,
		FoodID:     request.FoodID,
		Quantity:   request.Quantity,
	}

	return s.mealItemRepo.Create(newItem)
}

func (s *mealItemService) GetAll() ([]model.MealItem, error) {
	return s.mealItemRepo.FindAll()
}

func (s *mealItemService) GetByID(id uint) (*model.MealItem, error) {
	return s.mealItemRepo.FindByID(id)
}

func (s *mealItemService) Update(id uint, request request.UpdateMealItem) (*model.MealItem, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	item, err := s.mealItemRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	item.Date = request.Date
	item.MealTypeID = request.MealTypeID
	item.FoodID = request.FoodID
	item.Quantity = request.Quantity
	item.MealType = nil
	item.Food = nil
	return s.mealItemRepo.Update(item)
}

func (s *mealItemService) Delete(id uint) error {
	return s.mealItemRepo.Delete(id)
}
