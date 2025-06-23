package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type MealMenuRepo interface {
	Create(menu *model.MealMenu) (*model.MealMenu, error)
	FindAll() ([]model.MealMenu, error)
	FindByID(id uint) (*model.MealMenu, error)
	Update(menu *model.MealMenu) (*model.MealMenu, error)
	Delete(id uint) error
}

type mealMenuRepo struct {
	db *config.Database
}

func NewMealMenuRepo(db *config.Database) MealMenuRepo {
	return &mealMenuRepo{db: db}
}

func (r *mealMenuRepo) Create(menu *model.MealMenu) (*model.MealMenu, error) {
	tx := r.db.Gorm.Create(&menu)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var created model.MealMenu
	tx = r.db.Gorm.Preload("Foods").
		Preload("Foods.FoodMaterialUsages").
		Preload("Foods.FoodMaterialUsages.FoodMaterial").
		First(&created, menu.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &created, nil
}

func (r *mealMenuRepo) FindAll() ([]model.MealMenu, error) {
	var menus []model.MealMenu
	tx := r.db.Gorm.Preload("Foods").
		Preload("Foods.FoodMaterialUsages").
		Preload("Foods.FoodMaterialUsages.FoodMaterial").
		Preload("MealType").
		Find(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return menus, nil
}

func (r *mealMenuRepo) FindByID(id uint) (*model.MealMenu, error) {
	var menu model.MealMenu
	tx := r.db.Gorm.Preload("Foods").
		Preload("Foods.FoodMaterialUsages").
		Preload("Foods.FoodMaterialUsages.FoodMaterial").
		First(&menu, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &menu, nil
}

func (r *mealMenuRepo) Update(menu *model.MealMenu) (*model.MealMenu, error) {
	if err := r.db.Gorm.Model(menu).Association("Foods").Replace(menu.Foods); err != nil {
		return nil, err
	}
	tx := r.db.Gorm.Save(menu)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var updated model.MealMenu
	tx = r.db.Gorm.Preload("Foods").
		Preload("Foods.FoodMaterialUsages").
		Preload("Foods.FoodMaterialUsages.FoodMaterial").
		First(&updated, menu.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &updated, nil
}

func (r *mealMenuRepo) Delete(id uint) error {
	menu, err := r.FindByID(id)
	if err != nil {
		return err
	}
	err = r.db.Gorm.Model(&menu).Association("Foods").Clear()
	if err != nil {
		return err
	}
	tx := r.db.Gorm.Delete(&model.MealMenu{}, id)
	return tx.Error
}
