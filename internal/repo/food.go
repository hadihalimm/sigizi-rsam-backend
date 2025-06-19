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
	// for i := range food.FoodMaterialUsages {
	// 	food.FoodMaterialUsages[i].FoodID = food.ID
	// 	fmt.Printf("Inserting usage: %+v\n", food.FoodMaterialUsages[i])
	// 	if err := r.db.Gorm.Create(&food.FoodMaterialUsages[i]).Error; err != nil {
	// 		return nil, err
	// 	}
	// }
	var created model.Food
	tx = r.db.Gorm.Preload("FoodMaterialUsages.FoodMaterial").First(&created, food.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &created, nil
}

func (r *foodRepo) FindAll() ([]model.Food, error) {
	var foods []model.Food
	tx := r.db.Gorm.Preload("FoodMaterialUsages.FoodMaterial").Find(&foods)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return foods, nil
}

func (r *foodRepo) FindByID(id uint) (*model.Food, error) {
	var food model.Food
	tx := r.db.Gorm.Preload("FoodMaterialUsages.FoodMaterial").First(&food, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &food, nil
}

func (r *foodRepo) Update(food *model.Food) (*model.Food, error) {
	if err := r.db.Gorm.Where("food_id = ?", food.ID).Delete(&model.FoodMaterialUsage{}).Error; err != nil {
		return nil, err
	}
	if err := r.db.Gorm.Save(&food).Error; err != nil {
		return nil, err
	}

	var updated model.Food
	if err := r.db.Gorm.Preload("FoodMaterialUsages.FoodMaterial").First(&updated, food.ID).Error; err != nil {
		return nil, err
	}

	return &updated, nil

}

func (r *foodRepo) Delete(id uint) error {
	if err := r.db.Gorm.Where("food_id = ?", id).Delete(&model.FoodMaterialUsage{}).Error; err != nil {
		return err
	}
	if err := r.db.Gorm.Delete(&model.Food{}, id).Error; err != nil {
		return err
	}
	return nil
}
