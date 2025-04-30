package repo

import (
	"time"

	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type DailyPatientMealRepo interface {
	Create(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error)
	FindAll() ([]model.DailyPatientMeal, error)
	FindByID(id uint) (*model.DailyPatientMeal, error)
	Update(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error)
	Delete(id uint) error
	FilterByDateAndRoomType(
		date time.Time, roomType uint) ([]model.DailyPatientMeal, error)
	ReplaceDiets(meal *model.DailyPatientMeal, dietIDs []uint) error
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

	tx = r.db.Gorm.Preload("Diets").First(&meal, meal.ID)
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
	tx := r.db.Gorm.Preload("Diets").First(&meal, id)
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
	tx = r.db.Gorm.Preload("Diets").First(&meal, meal.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return meal, nil
}

func (r *dailyPatientMealRepo) Delete(id uint) error {
	var meal model.DailyPatientMeal
	tx := r.db.Gorm.Preload("Diets").First(&meal, id)
	if tx.Error != nil {
		return tx.Error
	}
	err := r.db.Gorm.Model(&meal).Association("Diets").Clear()
	if err != nil {
		return err
	}
	tx = r.db.Gorm.Delete(&model.DailyPatientMeal{}, id)
	return tx.Error
}

func (r *dailyPatientMealRepo) FilterByDateAndRoomType(
	date time.Time, roomType uint) ([]model.DailyPatientMeal, error) {

	var meals []model.DailyPatientMeal
	tx := r.db.Gorm.Preload("Patient").Preload("Room").
		Preload("Room.RoomType").Preload("MealType").Preload("Diets").
		Joins("JOIN rooms ON rooms.id = daily_patient_meals.room_id").
		Where("DATE(daily_patient_meals.created_at) = ?", date.Format("2006-01-02")).
		Where("rooms.room_type_id = ? ", roomType).
		Find(&meals)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return meals, nil
}

func (r *dailyPatientMealRepo) ReplaceDiets(meal *model.DailyPatientMeal, dietIDs []uint) error {
	var diets []model.Diet
	if len(dietIDs) == 0 {
		return r.db.Gorm.Model(&meal).Association("Diets").Clear()
	}
	tx := r.db.Gorm.Where("id IN ?", dietIDs).Find(&diets)
	if tx.Error != nil {
		return tx.Error
	}
	return r.db.Gorm.Model(&meal).Association("Diets").Replace(diets)
}
