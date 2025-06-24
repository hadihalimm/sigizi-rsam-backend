package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type MealMenuService interface {
	Create(request request.CreateMealMenu) (*model.MealMenu, error)
	FindAll() ([]model.MealMenu, error)
	FindByID(id uint) (*model.MealMenu, error)
	Update(id uint, request request.UpdateMealMenu) (*model.MealMenu, error)
	Delete(id uint) error
	CreateMealMenuTemplate(request request.CreateMealMenuTemplate) error
	FindAllMealMenuTemplate() ([]model.MealMenuTemplate, error)
	FindByIDMealMenuTemplate(id uint) (*model.MealMenuTemplate, error)
	UpdateMealMenuTemplate(
		id uint, request request.UpdateMealMenuTemplate) (*model.MealMenuTemplate, error)
	DeleteMealMenuTemplate(id uint) error
}

type mealMenuService struct {
	mealMenuRepo repo.MealMenuRepo
	foodRepo     repo.FoodRepo
	validate     *validator.Validate
}

func NewMealMenuService(
	mealMenuRepo repo.MealMenuRepo,
	foodRepo repo.FoodRepo,
	validate *validator.Validate) MealMenuService {
	return &mealMenuService{mealMenuRepo: mealMenuRepo, foodRepo: foodRepo, validate: validate}
}

func (s *mealMenuService) Create(request request.CreateMealMenu) (*model.MealMenu, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	foods, err := s.foodRepo.FindByIDs(request.FoodIDs)
	if err != nil {
		return nil, err
	}

	newMenu := &model.MealMenu{
		Name:               request.Name,
		Day:                request.Day,
		Time:               request.Time,
		MealTypeID:         request.MealTypeID,
		Foods:              foods,
		MealMenuTemplateID: request.MealMenuTemplateID,
	}

	return s.mealMenuRepo.Create(newMenu)
}

func (s *mealMenuService) FindAll() ([]model.MealMenu, error) {
	return s.mealMenuRepo.FindAll()
}

func (s *mealMenuService) FindByID(id uint) (*model.MealMenu, error) {
	return s.mealMenuRepo.FindByID(id)
}

func (s *mealMenuService) Update(id uint, request request.UpdateMealMenu) (*model.MealMenu, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	menu, err := s.mealMenuRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	foods, err := s.foodRepo.FindByIDs(request.FoodIDs)
	if err != nil {
		return nil, err
	}

	menu.Name = request.Name
	menu.Day = request.Day
	menu.Time = request.Time
	menu.MealTypeID = request.MealTypeID
	menu.Foods = foods

	return s.mealMenuRepo.Update(menu)
}

func (s *mealMenuService) Delete(id uint) error {
	return s.mealMenuRepo.Delete(id)
}

func (s *mealMenuService) CreateMealMenuTemplate(request request.CreateMealMenuTemplate) error {
	if err := s.validate.Struct(request); err != nil {
		return err
	}

	newTemplate := &model.MealMenuTemplate{
		Name: request.Name,
	}
	return s.mealMenuRepo.CreateNewMealMenuTemplate(newTemplate)
}

func (s *mealMenuService) FindAllMealMenuTemplate() ([]model.MealMenuTemplate, error) {
	return s.mealMenuRepo.FindAllMealMenuTemplate()
}

func (s *mealMenuService) FindByIDMealMenuTemplate(id uint) (*model.MealMenuTemplate, error) {
	return s.mealMenuRepo.FindByIDMealMenuTemplate(id)
}

func (s *mealMenuService) UpdateMealMenuTemplate(
	id uint, request request.UpdateMealMenuTemplate) (*model.MealMenuTemplate, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	template, err := s.mealMenuRepo.FindByIDMealMenuTemplate(id)
	if err != nil {
		return nil, err
	}
	template.Name = request.Name
	return s.mealMenuRepo.UpdateMealMenuTemplate(template)
}

func (s *mealMenuService) DeleteMealMenuTemplate(id uint) error {
	return s.mealMenuRepo.DeleteMealMenuTemplate(id)
}
