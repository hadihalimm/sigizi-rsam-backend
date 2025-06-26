package repo

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"gorm.io/gorm"
)

type DailyPatientMealRepo interface {
	Create(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error)
	FindAll() ([]model.DailyPatientMeal, error)
	FindByID(id uint) (*model.DailyPatientMeal, error)
	Update(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error)
	Delete(id uint) error
	FilterByDateAndRoomType(
		date time.Time, roomType uint) ([]model.DailyPatientMeal, error)
	FilterByDate(date time.Time) ([]model.DailyPatientMeal, error)
	InsertDiets(meal *model.DailyPatientMeal, dietIDs []uint) error
	ReplaceDiets(meal *model.DailyPatientMeal, dietIDs []uint) error
	CountByDateAndRoomType(
		date time.Time, roomTypeID uint) ([]MealMatrixEntry, error)
	CountByDateForAllRoomTypes(
		date time.Time) ([]MealMatrixEntry, error)
	CountForEveryMealType(date time.Time) ([]DailyMealTypeCount, error)
	FilterLogsByDate(date time.Time) ([]model.DailyPatientMealLog, error)
	FilterLogsByDateAndRoomType(
		date time.Time, roomTypeID uint) ([]model.DailyPatientMealLog, error)
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
	tx := r.db.Gorm.Preload("Diets").Preload("Room").Preload("Room.RoomType").First(&meal, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &meal, nil
}

func (r *dailyPatientMealRepo) Update(meal *model.DailyPatientMeal) (*model.DailyPatientMeal, error) {
	var existingMeal model.DailyPatientMeal
	var reloadedMeal model.DailyPatientMeal
	err := r.db.Gorm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Room").
			Preload("Patient").
			Preload("MealType").
			Preload("Room.RoomType").
			First(&existingMeal, meal.ID).Error; err != nil {
			return err
		}
		if err := tx.Save(meal).Error; err != nil {
			return err
		}
		var updatedMeal model.DailyPatientMeal
		if err := tx.Preload("Room").
			Preload("Patient").
			Preload("MealType").
			Preload("Room.RoomType").
			First(&updatedMeal, meal.ID).Error; err != nil {
			return err
		}

		var logs []model.DailyPatientMealLog
		addLog := func(field string, oldVal, newVal interface{}) {
			oldStr := fmt.Sprintf("%v", oldVal)
			newStr := fmt.Sprintf("%v", newVal)
			if oldStr != newStr {
				logs = append(logs, model.DailyPatientMealLog{
					DailyPatientMealID: updatedMeal.ID,
					RoomTypeName:       updatedMeal.Room.RoomType.Name,
					RoomName:           updatedMeal.Room.Name,
					PatientMRN:         updatedMeal.Patient.MedicalRecordNumber,
					PatientName:        updatedMeal.Patient.Name,
					Field:              field,
					OldValue:           oldStr,
					NewValue:           newStr,
					ChangedAt:          time.Now(),
					Date:               updatedMeal.Date.Truncate((24 * time.Hour)),
				})
			}
		}
		addLog("RoomID", existingMeal.Room.Name, updatedMeal.Room.Name)
		addLog("RoomType", existingMeal.Room.RoomType.Name, updatedMeal.Room.RoomType.Name)
		addLog("MealTypeID", existingMeal.MealTypeID, updatedMeal.MealTypeID)
		addLog("Notes", existingMeal.Notes, updatedMeal.Notes)
		if len(logs) > 0 {
			if err := tx.Create(&logs).Error; err != nil {
				return err
			}
		}
		if err := tx.Preload("Room").
			Preload("Room.RoomType").
			Preload("Patient").
			Preload("MealType").
			Preload("Diets").
			First(&reloadedMeal, meal.ID).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &reloadedMeal, err
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

func (r *dailyPatientMealRepo) InsertDiets(meal *model.DailyPatientMeal, dietIDs []uint) error {
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

func (r *dailyPatientMealRepo) ReplaceDiets(meal *model.DailyPatientMeal, dietIDs []uint) error {
	return r.db.Gorm.Transaction(func(tx *gorm.DB) error {
		var currentDiets []model.Diet
		if err := tx.Model(meal).Association("Diets").Find(&currentDiets); err != nil {
			return err
		}

		oldDietIDs := []string{}
		for _, d := range currentDiets {
			oldDietIDs = append(oldDietIDs, fmt.Sprintf("%d", d.ID))
		}
		oldDietStr := strings.Join(oldDietIDs, ",")

		var newDiets []model.Diet
		var newDietStr string

		if len(dietIDs) == 0 {
			if err := tx.Model(meal).Association("Diets").Clear(); err != nil {
				return err
			}
			newDietStr = ""
		} else {
			if err := tx.Where("id IN ?", dietIDs).Find(&newDiets).Error; err != nil {
				return err
			}
			sort.Slice(newDiets, func(i, j int) bool {
				return newDiets[i].ID < newDiets[j].ID
			})
			if err := tx.Model(meal).Association("Diets").Replace(newDiets); err != nil {
				return err
			}
			newDietIDs := []string{}
			for _, d := range newDiets {
				newDietIDs = append(newDietIDs, fmt.Sprintf("%d", d.ID))
			}
			newDietStr = strings.Join(newDietIDs, ",")
		}
		if oldDietStr != newDietStr {
			log := model.DailyPatientMealLog{
				DailyPatientMealID: meal.ID,
				RoomTypeName:       meal.Room.RoomType.Name,
				RoomName:           meal.Room.Name,
				PatientMRN:         meal.Patient.MedicalRecordNumber,
				PatientName:        meal.Patient.Name,
				Field:              "Diets",
				OldValue:           oldDietStr,
				NewValue:           newDietStr,
				ChangedAt:          time.Now(),
				Date:               meal.Date.Truncate((24 * time.Hour)),
			}
			if err := tx.Create(&log).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *dailyPatientMealRepo) FilterByDateAndRoomType(
	date time.Time, roomType uint) ([]model.DailyPatientMeal, error) {

	var meals []model.DailyPatientMeal
	tx := r.db.Gorm.Preload("Patient").Preload("Patient.Allergies").Preload("Room").
		Preload("Room.RoomType").Preload("MealType").Preload("Diets").
		Joins("JOIN rooms ON rooms.id = daily_patient_meals.room_id").
		Where("DATE(daily_patient_meals.date) = ?", date.Format("2006-01-02")).
		Where("rooms.room_type_id = ? ", roomType).
		Find(&meals)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return meals, nil
}

func (r *dailyPatientMealRepo) FilterByDate(date time.Time) ([]model.DailyPatientMeal, error) {
	var meals []model.DailyPatientMeal
	tx := r.db.Gorm.Preload("Patient").Preload("Patient.Allergies").Preload("Room").
		Preload("Room.RoomType").Preload("MealType").Preload("Diets").
		Joins("JOIN rooms ON rooms.id = daily_patient_meals.room_id").
		Joins("JOIN room_types ON room_types.id = rooms.room_type_id").
		Where("DATE(daily_patient_meals.date) = ?", date.Format("2006-01-02")).
		Order("room_types.id").
		Find(&meals)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return meals, nil
}

type MealMatrixEntry struct {
	TreatmentClass string `json:"treatmentClass"`
	MealType       string `json:"mealType"`
	MealCount      string `json:"mealCount"`
}

func (r *dailyPatientMealRepo) CountByDateAndRoomType(
	date time.Time, roomTypeID uint) ([]MealMatrixEntry, error) {

	var results []MealMatrixEntry
	tx := r.db.Gorm.Table("daily_patient_meals").
		Select("rooms.treatment_class AS treatment_class, meal_types.code AS meal_type, COUNT(daily_patient_meals.id) AS meal_count").
		Joins("JOIN rooms ON rooms.id = daily_patient_meals.room_id").
		Joins("JOIN meal_types ON meal_types.id = daily_patient_meals.meal_type_id").
		Where("rooms.room_type_id = ?", roomTypeID).
		Where("DATE(daily_patient_meals.date) = ?", date.Format("2006-01-02")).
		Group("rooms.treatment_class, meal_types.name").
		Order("rooms.treatment_class, meal_types.name").
		Scan(&results)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return results, nil
}

func (r *dailyPatientMealRepo) CountByDateForAllRoomTypes(
	date time.Time) ([]MealMatrixEntry, error) {

	var results []MealMatrixEntry
	tx := r.db.Gorm.Table("daily_patient_meals").
		Select("rooms.treatment_class AS treatment_class, meal_types.code AS meal_type, COUNT(daily_patient_meals.id) AS meal_count").
		Joins("JOIN rooms ON rooms.id = daily_patient_meals.room_id").
		Joins("JOIN meal_types ON meal_types.id = daily_patient_meals.meal_type_id").
		Where("DATE(daily_patient_meals.date) = ?", date.Format("2006-01-02")).
		Group("rooms.treatment_class, meal_types.name").
		Order("rooms.treatment_class, meal_types.name").
		Scan(&results)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return results, nil
}

type DailyMealTypeCount struct {
	MealTypeID   uint   `json:"mealTypeID"`
	MealTypeCode string `json:"mealTypeCode"`
	MealCount    uint   `json:"mealCount"`
}

func (r *dailyPatientMealRepo) CountForEveryMealType(date time.Time) ([]DailyMealTypeCount, error) {
	var results []DailyMealTypeCount
	tx := r.db.Gorm.Table("daily_patient_meals").
		Select("meal_types.id AS meal_type_id, meal_types.code AS meal_type_code, COUNT(daily_patient_meals.id) AS meal_count").
		Joins("JOIN meal_types ON meal_types.id = daily_patient_meals.meal_type_id").
		Where("DATE(daily_patient_meals.date) = ?", date.Format("2006-01-02")).
		Group("meal_types.id, meal_types.name").
		Scan(&results)

	if tx.Error != nil {
		return nil, tx.Error
	}
	if results == nil {
		return []DailyMealTypeCount{}, nil
	}
	return results, nil
}

func (r *dailyPatientMealRepo) FilterLogsByDate(date time.Time) ([]model.DailyPatientMealLog, error) {
	var logs []model.DailyPatientMealLog
	err := r.db.Gorm.
		Where("DATE(date) = ?", date.Format("2006-01-02")).
		Order("changed_at DESC").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *dailyPatientMealRepo) FilterLogsByDateAndRoomType(
	date time.Time, roomTypeID uint) ([]model.DailyPatientMealLog, error) {
	var logs []model.DailyPatientMealLog
	err := r.db.Gorm.
		Where("DATE(date) = ?", date.Format("2006-01-02")).
		Where("room_type_id = ? ", roomTypeID).
		Order("changed_at DESC").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}
