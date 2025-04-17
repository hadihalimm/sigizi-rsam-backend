package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type DailyPatientMealRepo interface {
	Create(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error)
	FindAll() ([]model.DailyPatientMeal, error)
	FindByID(id uint) (*model.DailyPatientMeal, error)
	Update(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error)
	Delete(id uint) error
}

type dailyPatientMealRepo struct {
	db *config.Database
}

func NewDailyPatientMealRepo(db *config.Database) DailyPatientMealRepo {
	return &dailyPatientMealRepo{db: db}
}

func (r *dailyPatientMealRepo) Create(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error) {
	tx := r.db.Gorm.Create(&meal)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = r.db.Gorm.First(&meal, meal.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return meal, nil
}

func (r *dailyPatientMealRepo) FindAll() ([]model.DailyPatientMeal, error) {
	var meals []model.DailyPatientMeal
	tx := r.db.Gorm.Find(&meals)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return meals, nil
}

func (r *dailyPatientMealRepo) FindByID(id uint) (*model.DailyPatientMeal, error) {
	var meal model.DailyPatientMeal
	tx := r.db.Gorm.First(&meal, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &meal, nil
}

func (r *dailyPatientMealRepo) Update(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error) {
	tx := r.db.Gorm.Save(meal)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return meal, nil
}

func (r *dailyPatientMealRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.DailyPatientMeal{}, id)
	return tx.Error
}
