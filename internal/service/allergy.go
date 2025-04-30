package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type AllergyService interface {
	Create(request request.CreateAllergy) (*model.Allergy, error)
	GetAll() ([]model.Allergy, error)
	GetByID(id uint) (*model.Allergy, error)
	Update(id uint, request request.UpdateAllergy) (*model.Allergy, error)
	Delete(id uint) error
}

type allergyService struct {
	allergyRepo repo.AllergyRepo
	validate    *validator.Validate
}

func NewAllergyService(allergyRepo repo.AllergyRepo, validate *validator.Validate) AllergyService {
	return &allergyService{allergyRepo: allergyRepo, validate: validate}
}

func (s *allergyService) Create(request request.CreateAllergy) (*model.Allergy, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newAllergy := &model.Allergy{
		Code: request.Code,
		Name: request.Name,
	}
	return s.allergyRepo.Create(newAllergy)
}

func (s *allergyService) GetAll() ([]model.Allergy, error) {
	return s.allergyRepo.FindAll()
}

func (s *allergyService) GetByID(id uint) (*model.Allergy, error) {
	return s.allergyRepo.FindByID(id)
}

func (s *allergyService) Update(id uint, request request.UpdateAllergy) (*model.Allergy, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	allergy, err := s.allergyRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("allergy not found")
	}

	allergy.Code = request.Code
	allergy.Name = request.Name
	return s.allergyRepo.Update(allergy)
}

func (s *allergyService) Delete(id uint) error {
	return s.allergyRepo.Delete(id)
}
