package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type FoodMaterialRepo interface {
	Create(food *model.FoodMaterial) (*model.FoodMaterial, error)
	FindAll() ([]model.FoodMaterial, error)
	FindByID(id uint) (*model.FoodMaterial, error)
	Update(food *model.FoodMaterial) (*model.FoodMaterial, error)
	Delete(id uint) error
}

type foodMaterialRepo struct {
	db *config.Database
}

func NewFoodRepo(db *config.Database) FoodMaterialRepo {
	return &foodMaterialRepo{db: db}
}

func (r *foodMaterialRepo) Create(foodMaterial *model.FoodMaterial) (*model.FoodMaterial, error) {
	tx := r.db.Gorm.Create(&foodMaterial)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return foodMaterial, nil
}

func (r *foodMaterialRepo) FindAll() ([]model.FoodMaterial, error) {
	var foodMaterials []model.FoodMaterial
	tx := r.db.Gorm.Find(&foodMaterials)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return foodMaterials, nil
}

func (r *foodMaterialRepo) FindByID(id uint) (*model.FoodMaterial, error) {
	var foodMaterial model.FoodMaterial
	tx := r.db.Gorm.First(&foodMaterial, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &foodMaterial, nil
}

func (r *foodMaterialRepo) Update(foodMaterial *model.FoodMaterial) (*model.FoodMaterial, error) {
	tx := r.db.Gorm.Save(foodMaterial)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return foodMaterial, nil
}

func (r *foodMaterialRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.FoodMaterial{}, id)
	return tx.Error
}
