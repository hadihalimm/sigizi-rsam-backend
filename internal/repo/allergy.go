package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type AllergyRepo interface {
	Create(allergy *model.Allergy) (*model.Allergy, error)
	FindAll() ([]model.Allergy, error)
	FindByID(id uint) (*model.Allergy, error)
	Update(allergy *model.Allergy) (*model.Allergy, error)
	Delete(id uint) error
}

type allergyRepo struct {
	db *config.Database
}

func NewAllergyRepo(db *config.Database) AllergyRepo {
	return &allergyRepo{db: db}
}

func (r *allergyRepo) Create(allergy *model.Allergy) (*model.Allergy, error) {
	tx := r.db.Gorm.Create(&allergy)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = r.db.Gorm.First(&allergy, allergy.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return allergy, nil
}

func (r *allergyRepo) FindAll() ([]model.Allergy, error) {
	var allergies []model.Allergy
	tx := r.db.Gorm.Find(&allergies)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return allergies, nil
}

func (r *allergyRepo) FindByID(id uint) (*model.Allergy, error) {
	var allergy model.Allergy
	tx := r.db.Gorm.First(&allergy, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &allergy, nil
}

func (r *allergyRepo) Update(allergy *model.Allergy) (*model.Allergy, error) {
	tx := r.db.Gorm.Save(allergy)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return allergy, nil
}

func (r *allergyRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.Allergy{}, id)
	return tx.Error
}
