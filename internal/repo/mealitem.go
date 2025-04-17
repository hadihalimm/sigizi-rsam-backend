package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type MealItemRepo interface {
	Create(item *model.MealItem) (*model.MealItem, error)
	FindAll() ([]model.MealItem, error)
	FindByID(id uint) (*model.MealItem, error)
	Update(item *model.MealItem) (*model.MealItem, error)
	Delete(id uint) error
}

type mealItemRepo struct {
	db *config.Database
}

func NewMealItemRepo(db *config.Database) MealItemRepo {
	return &mealItemRepo{db: db}
}

func (r *mealItemRepo) Create(item *model.MealItem) (*model.MealItem, error) {
	tx := r.db.Gorm.Create(&item)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = r.db.Gorm.Preload("MealType").Preload("Food").First(&item, item.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return item, nil
}

func (r *mealItemRepo) FindAll() ([]model.MealItem, error) {
	var items []model.MealItem
	tx := r.db.Gorm.Preload("MealType").Preload("Food").Find(&items)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return items, nil
}

func (r *mealItemRepo) FindByID(id uint) (*model.MealItem, error) {
	var item model.MealItem
	tx := r.db.Gorm.Preload("MealType").Preload("Food").First(&item, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &item, nil
}

func (r *mealItemRepo) Update(item *model.MealItem) (*model.MealItem, error) {
	tx := r.db.Gorm.Save(item)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := r.db.Gorm.Preload("MealType").Preload("Food").First(&item, item.ID).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (r *mealItemRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.MealItem{}, id)
	return tx.Error
}
