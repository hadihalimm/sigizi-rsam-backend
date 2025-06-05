package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
	"github.com/xuri/excelize/v2"
)

type DailyPatientMealService interface {
	Create(request request.CreateDailyPatientMeal) (*model.DailyPatientMeal, error)
	GetAll() ([]model.DailyPatientMeal, error)
	GetByID(id uint) (*model.DailyPatientMeal, error)
	Update(id uint, request request.UpdateDailyPatientMeal) (*model.DailyPatientMeal, error)
	Delete(id uint) error
	FilterByDateAndRoomType(
		date time.Time, roomTypeID uint) ([]model.DailyPatientMeal, error)
	CountByDateAndRoomType(
		date time.Time, roomTypeID uint) ([]repo.MealMatrixEntry, error)
	ExportToExcel(date time.Time) (*excelize.File, error)
	FilterLogsByDate(date time.Time) ([]model.DailyPatientMealLog, error)
	FilterLogsByDateAndRoomType(
		date time.Time, roomTypeID uint) ([]model.DailyPatientMealLog, error)
	CountDietCombinationsByDate(
		date time.Time) ([]DietCombinationCount, int, int, error)
}

type dailyPatientMealService struct {
	dailyPatientMealRepo repo.DailyPatientMealRepo
	roomTypeRepo         repo.RoomTypeRepo
	validate             *validator.Validate
}

func NewDailyPatientMealService(
	dailyPatientMealRepo repo.DailyPatientMealRepo,
	roomTypeRepo repo.RoomTypeRepo,
	validate *validator.Validate,
) DailyPatientMealService {
	return &dailyPatientMealService{
		dailyPatientMealRepo: dailyPatientMealRepo,
		roomTypeRepo:         roomTypeRepo,
		validate:             validate}
}

func (s *dailyPatientMealService) Create(request request.CreateDailyPatientMeal) (*model.DailyPatientMeal, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newDailyMeal := &model.DailyPatientMeal{
		PatientID:       request.PatientID,
		RoomID:          request.RoomID,
		MealTypeID:      request.MealTypeID,
		Date:            request.Date.Truncate((24 * time.Hour)),
		Notes:           request.Notes,
		IsNewlyAdmitted: request.IsNewlyAdmitted,
	}

	meal, err := s.dailyPatientMealRepo.Create(newDailyMeal)
	if err != nil {
		return nil, err
	}

	err = s.dailyPatientMealRepo.InsertDiets(meal, request.DietIDs)
	if err != nil {
		return nil, err
	}
	meal, err = s.dailyPatientMealRepo.FindByID(meal.ID)
	if err != nil {
		return nil, err
	}
	return meal, nil
}

func (s *dailyPatientMealService) GetAll() ([]model.DailyPatientMeal, error) {
	return s.dailyPatientMealRepo.FindAll()
}

func (s *dailyPatientMealService) GetByID(id uint) (*model.DailyPatientMeal, error) {
	return s.dailyPatientMealRepo.FindByID(id)
}

func (s *dailyPatientMealService) Update(id uint, request request.UpdateDailyPatientMeal) (*model.DailyPatientMeal, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}
	meal, err := s.dailyPatientMealRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("daily patient meal not found")
	}

	meal.PatientID = request.PatientID
	meal.Room = model.Room{}
	meal.RoomID = request.RoomID
	meal.MealTypeID = request.MealTypeID
	meal.Notes = request.Notes
	meal.IsNewlyAdmitted = request.IsNewlyAdmitted

	meal, err = s.dailyPatientMealRepo.Update(meal)
	if err != nil {
		return nil, err
	}
	err = s.dailyPatientMealRepo.ReplaceDiets(meal, request.DietIDs)
	if err != nil {
		return nil, err
	}
	meal, err = s.dailyPatientMealRepo.FindByID(meal.ID)
	if err != nil {
		return nil, err
	}
	return meal, nil
}

func (s *dailyPatientMealService) Delete(id uint) error {
	return s.dailyPatientMealRepo.Delete(id)
}

func (s *dailyPatientMealService) FilterByDateAndRoomType(
	date time.Time, roomTypeID uint) ([]model.DailyPatientMeal, error) {
	if roomTypeID == 0 {
		return s.FilterByDate(date)
	}
	_, err := s.roomTypeRepo.FindByID(roomTypeID)
	if err != nil {
		return nil, errors.New("room type not found")
	}
	return s.dailyPatientMealRepo.FilterByDateAndRoomType(date, roomTypeID)
}

func (s *dailyPatientMealService) FilterByDate(date time.Time) ([]model.DailyPatientMeal, error) {
	return s.dailyPatientMealRepo.FilterByDate(date)
}

func (s *dailyPatientMealService) CountByDateAndRoomType(
	date time.Time, roomTypeID uint) ([]repo.MealMatrixEntry, error) {
	if roomTypeID == 0 {
		return s.dailyPatientMealRepo.CountByDateForAllRoomTypes(date)
	}

	_, err := s.roomTypeRepo.FindByID(roomTypeID)
	if err != nil {
		return nil, errors.New("room type not found")
	}
	return s.dailyPatientMealRepo.CountByDateAndRoomType(date, roomTypeID)
}

func (s *dailyPatientMealService) ExportToExcel(date time.Time) (*excelize.File, error) {
	meals, err := s.FilterByDate(date)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Permintaan Makanan"
	f.NewSheet(sheet)

	headers := []string{"ID", "Nama Pasien", "No. MR", "Tanggal Lahir", "Kamar", "Diet", "Catatan"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, m := range meals {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), m.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), m.Patient.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), m.Patient.MedicalRecordNumber)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), m.Patient.DateOfBirth.Format("02-01-2006"))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), fmt.Sprintf("%s - %s", m.Room.RoomType.Name, m.Room.Name))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), fmt.Sprintf("%s%s %s", m.MealType.Code,
			strings.Join(extractDietCodes(m.Diets), ""),
			strings.Join(extractAllergyCodes(m.Patient.Allergies), "")))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), m.Notes)

		f.SetColWidth(sheet, "A", "A", 5)
		f.SetColWidth(sheet, "B", "B", 20)
		f.SetColWidth(sheet, "C", "C", 10)
		f.SetColWidth(sheet, "D", "D", 12)
		f.SetColWidth(sheet, "E", "E", 20)
		f.SetColWidth(sheet, "F", "F", 30)
		f.SetColWidth(sheet, "G", "G", 20)
	}

	f.DeleteSheet("Sheet1")
	return f, nil
}

func extractDietCodes(diets []model.Diet) []string {
	names := make([]string, len(diets))
	for i, d := range diets {
		names[i] = d.Code
	}
	return names
}

func extractAllergyCodes(allergies []model.Allergy) []string {
	names := make([]string, len(allergies))
	for i, a := range allergies {
		names[i] = a.Code
	}
	return names
}

func (s *dailyPatientMealService) FilterLogsByDateAndRoomType(
	date time.Time, roomTypeID uint) ([]model.DailyPatientMealLog, error) {
	return s.dailyPatientMealRepo.FilterLogsByDateAndRoomType(date, roomTypeID)
}

func (s *dailyPatientMealService) FilterLogsByDate(date time.Time) ([]model.DailyPatientMealLog, error) {
	return s.dailyPatientMealRepo.FilterLogsByDate(date)
}

type DietCombinationCount struct {
	DietCodes string `json:"dietCodes"` // e.g., "DIAB-HYPO"
	Count     int    `json:"count"`     // number of meals with this combination
}

func (s *dailyPatientMealService) CountDietCombinationsByDate(
	date time.Time) ([]DietCombinationCount, int, int, error) {
	meals, err := s.dailyPatientMealRepo.FilterByDate(date)
	if err != nil {
		return nil, 0, 0, err
	}

	countMap := make(map[string]int)
	complicationCount := 0
	nonComplicationCount := 0
	for _, meal := range meals {
		if len(meal.Diets) == 0 {
			continue
		}
		if len(meal.Diets) == 1 {
			nonComplicationCount++
		} else {
			complicationCount++
		}
		var keyParts []string
		for _, diet := range meal.Diets {
			keyParts = append(keyParts, diet.Code)
		}
		key := strings.Join(keyParts, "")
		countMap[key]++
	}

	var result []DietCombinationCount
	for key, count := range countMap {
		result = append(result, DietCombinationCount{
			DietCodes: key,
			Count:     count,
		})
	}
	return result, complicationCount, nonComplicationCount, nil
}
