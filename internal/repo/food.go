package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type FoodRepo interface {
	Create(item *model.Food) (*model.Food, error)
	FindAll() ([]model.Food, error)
	FindByID(id uint) (*model.Food, error)
	Update(item *model.Food) (*model.Food, error)
	Delete(id uint) error
}

type foodRepo struct {
	db *config.Database
}

func NewFoodRepo(db *config.Database) FoodRepo {
	return &foodRepo{db: db}
}

func (r *foodRepo) Create(food *model.Food) (*model.Food, error) {
	tx := r.db.Gorm.Create(&food)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = r.db.Gorm.Preload("MealType").Preload("Food").First(&food, food.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return food, nil
}

func (r *foodRepo) FindAll() ([]model.Food, error) {
	var foods []model.Food
	tx := r.db.Gorm.Preload("MealType").Preload("Food").Find(&foods)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return foods, nil
}

func (r *foodRepo) FindByID(id uint) (*model.Food, error) {
	var food model.Food
	tx := r.db.Gorm.Preload("MealType").Preload("Food").First(&food, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &food, nil
}

func (r *foodRepo) Update(food *model.Food) (*model.Food, error) {
	tx := r.db.Gorm.Save(food)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := r.db.Gorm.Preload("MealType").Preload("Food").First(&food, food.ID).Error; err != nil {
		return nil, err
	}

	return food, nil
}

func (r *foodRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.Food{}, id)
	return tx.Error
}
