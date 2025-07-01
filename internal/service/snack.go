package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type SnackService interface {
	Create(request request.CreateSnack) (*model.Snack, error)
	FindAll() ([]model.Snack, error)
	FindByID(id uint) (*model.Snack, error)
	Update(id uint, request request.UpdateSnack) (*model.Snack, error)
	Delete(id uint) error

	CreateVariant(
		snackID uint, request request.CreateSnackVariant) (*model.SnackVariant, error)
	FindAllVariant(snackID uint) ([]model.SnackVariant, error)
	FindVariantByID(id uint) (*model.SnackVariant, error)
	UpdateVariant(
		id uint, request request.UpdateSnackVariant) (*model.SnackVariant, error)
	DeleteVariant(id uint) error
}

type snackService struct {
	snackRepo repo.SnackRepo
	validate  *validator.Validate
}

func NewSnackService(snackRepo repo.SnackRepo, validate *validator.Validate) SnackService {
	return &snackService{snackRepo: snackRepo, validate: validate}
}

func (s *snackService) Create(request request.CreateSnack) (*model.Snack, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newSnack := &model.Snack{
		Name:          request.Name,
		SnackVariants: []model.SnackVariant{},
	}
	return s.snackRepo.Create(newSnack)
}

func (s *snackService) FindAll() ([]model.Snack, error) {
	return s.snackRepo.FindAll()
}

func (s *snackService) FindByID(id uint) (*model.Snack, error) {
	return s.snackRepo.FindByID(id)
}

func (s *snackService) Update(id uint, request request.UpdateSnack) (*model.Snack, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	snack, err := s.snackRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	snack.Name = request.Name
	updatedSnack, err := s.snackRepo.Update(snack)
	if err != nil {
		return nil, err
	}
	return updatedSnack, nil
}

func (s *snackService) Delete(id uint) error {
	return s.snackRepo.Delete(id)
}

func (s *snackService) CreateVariant(
	snackID uint, request request.CreateSnackVariant) (*model.SnackVariant, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	var usages []model.SnackVariantMaterialUsage
	for _, usage := range request.SnackVariantMaterialUsages {
		usages = append(usages, model.SnackVariantMaterialUsage{
			FoodMaterialID: usage.FoodMaterialID,
			QuantityUsed:   usage.QuantityUsed,
		})
	}

	newVariant := &model.SnackVariant{
		SnackID:                    snackID,
		Name:                       request.Name,
		SnackVariantMaterialUsages: usages,
	}

	updatedVariant, err := s.snackRepo.CreateVariant(newVariant)
	if err != nil {
		return nil, err
	}
	if err := s.snackRepo.ReplaceDiets(updatedVariant, request.DietIDs); err != nil {
		return nil, err
	}
	if err := s.snackRepo.ReplaceMealTypes(updatedVariant, request.MealTypeIDs); err != nil {
		return nil, err
	}
	return s.snackRepo.FindVariantByID(updatedVariant.ID)
}

func (s *snackService) FindAllVariant(snackID uint) ([]model.SnackVariant, error) {
	return s.snackRepo.FindAllVariant(snackID)
}

func (s *snackService) FindVariantByID(id uint) (*model.SnackVariant, error) {
	return s.snackRepo.FindVariantByID(id)
}

func (s *snackService) UpdateVariant(
	id uint, request request.UpdateSnackVariant) (*model.SnackVariant, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	variant, err := s.snackRepo.FindVariantByID(id)
	if err != nil {
		return nil, err
	}
	var usages []model.SnackVariantMaterialUsage
	for _, usage := range request.SnackVariantMaterialUsages {
		usages = append(usages, model.SnackVariantMaterialUsage{
			FoodMaterialID: usage.FoodMaterialID,
			QuantityUsed:   usage.QuantityUsed,
		})
	}

	variant.Name = request.Name
	variant.SnackVariantMaterialUsages = usages
	updatedVariant, err := s.snackRepo.UpdateVariant(variant)
	if err != nil {
		return nil, err
	}
	if err := s.snackRepo.ReplaceDiets(updatedVariant, request.DietIDs); err != nil {
		return nil, err
	}
	if err := s.snackRepo.ReplaceMealTypes(updatedVariant, request.MealTypeIDs); err != nil {
		return nil, err
	}
	return s.FindVariantByID(id)
}

func (s *snackService) DeleteVariant(id uint) error {
	variant, err := s.snackRepo.FindVariantByID(id)
	if err != nil {
		return err
	}
	if err := s.snackRepo.ReplaceDiets(variant, []uint{}); err != nil {
		return err
	}
	if err := s.snackRepo.ReplaceMealTypes(variant, []uint{}); err != nil {
		return err
	}
	return s.snackRepo.DeleteVariant(id)
}
