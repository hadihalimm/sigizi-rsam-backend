package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type DietRepo interface {
	Create(diet *model.Diet) (*model.Diet, error)
	FindAll() ([]model.Diet, error)
	FindByID(id uint) (*model.Diet, error)
	Update(diet *model.Diet) (*model.Diet, error)
	Delete(id uint) error
}

type dietRepo struct {
	db *config.Database
}

func NewDietRepo(db *config.Database) DietRepo {
	return &dietRepo{db: db}
}

func (r *dietRepo) Create(diet *model.Diet) (*model.Diet, error) {
	tx := r.db.Gorm.Create(&diet)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = r.db.Gorm.First(&diet, diet.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return diet, nil
}

func (r *dietRepo) FindAll() ([]model.Diet, error) {
	var diets []model.Diet
	tx := r.db.Gorm.Find(&diets)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return diets, nil
}

func (r *dietRepo) FindByID(id uint) (*model.Diet, error) {
	var diet model.Diet
	tx := r.db.Gorm.First(&diet, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &diet, nil
}

func (r *dietRepo) Update(diet *model.Diet) (*model.Diet, error) {
	tx := r.db.Gorm.Save(diet)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return diet, nil
}

func (r *dietRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.Diet{}, id)
	return tx.Error
}
