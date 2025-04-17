package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type DailyPatientMealService interface {
	Create(request request.CreateDailyPatientMeal) (*model.DailyPatientMeal, error)
	GetAll() ([]model.DailyPatientMeal, error)
	GetByID(id uint) (*model.DailyPatientMeal, error)
	Update(id uint, request request.UpdateDailyPatientMeal) (*model.DailyPatientMeal, error)
	Delete(id uint) error
}

type dailyPatientMealService struct {
	dailyPatientMealRepo repo.DailyPatientMealRepo
	validate             *validator.Validate
}

func NewDailyPatientMealService(
	dailyPatientMealRepo repo.DailyPatientMealRepo,
	validate *validator.Validate,
) DailyPatientMealService {
	return &dailyPatientMealService{dailyPatientMealRepo: dailyPatientMealRepo, validate: validate}
}

func (s *dailyPatientMealService) Create(request request.CreateDailyPatientMeal) (*model.DailyPatientMeal, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newDailyMeal := &model.DailyPatientMeal{
		Date:       request.Date,
		PatientID:  request.PatientID,
		RoomID:     request.RoomID,
		MealTypeID: request.MealTypeID,
		Notes:      request.Notes,
	}
	return s.dailyPatientMealRepo.Create(newDailyMeal)
}

func (s *dailyPatientMealService) GetAll() ([]model.DailyPatientMeal, error) {
	return s.dailyPatientMealRepo.FindAll()
}

func (s *dailyPatientMealService) GetByID(id uint) (*model.DailyPatientMeal, error) {
	return s.dailyPatientMealRepo.FindByID(id)
}

func (s *dailyPatientMealService) Update(id uint, request request.UpdateDailyPatientMeal) (*model.DailyPatientMeal, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	meal, err := s.dailyPatientMealRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("daily patient meal not found")
	}

	meal.Date = request.Date
	meal.PatientID = request.PatientID
	meal.RoomID = request.RoomID
	meal.MealTypeID = request.MealTypeID
	meal.Notes = request.Notes
	return s.dailyPatientMealRepo.Update(meal)
}

func (s *dailyPatientMealService) Delete(id uint) error {
	return s.dailyPatientMealRepo.Delete(id)
}
