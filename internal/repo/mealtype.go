package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type MealTypeRepo interface {
	Create(mt *model.MealType) (*model.MealType, error)
	FindAll() ([]model.MealType, error)
	FindByID(id uint) (*model.MealType, error)
	Update(mt *model.MealType) (*model.MealType, error)
	Delete(id uint) error
}

type mealTypeRepo struct {
	db *config.Database
}

func NewMealTypeRepo(db *config.Database) MealTypeRepo {
	return &mealTypeRepo{db: db}
}

func (r *mealTypeRepo) Create(mt *model.MealType) (*model.MealType, error) {
	tx := r.db.Gorm.Create(&mt)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mt, nil
}

func (r *mealTypeRepo) FindAll() ([]model.MealType, error) {
	var mealTypes []model.MealType
	tx := r.db.Gorm.Find(&mealTypes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mealTypes, nil
}

func (r *mealTypeRepo) FindByID(id uint) (*model.MealType, error) {
	var mt model.MealType
	tx := r.db.Gorm.First(&mt, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &mt, nil
}

func (r *mealTypeRepo) Update(mt *model.MealType) (*model.MealType, error) {
	tx := r.db.Gorm.Save(mt)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mt, nil
}

func (r *mealTypeRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.MealType{}, id)
	return tx.Error
}
