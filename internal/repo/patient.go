package repo

import (
	"strings"

	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type PatientRepo interface {
	Create(patient *model.Patient) (*model.Patient, error)
	FindAll() ([]model.Patient, error)
	FindByID(id uint) (*model.Patient, error)
	Update(patient *model.Patient) (*model.Patient, error)
	Delete(id uint) error
	FilterByMRN(mrn string) (*model.Patient, error)
	FindAllWithPaginationAndKeyword(limit int, offset int, keyword string) ([]model.Patient, int64, error)
	ReplaceAllergies(patient *model.Patient, allergyIDs []uint) error
}

type patientRepo struct {
	db *config.Database
}

func NewPatientRepo(db *config.Database) PatientRepo {
	return &patientRepo{db: db}
}

func (r *patientRepo) Create(patient *model.Patient) (*model.Patient, error) {
	tx := r.db.Gorm.Create(&patient)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return patient, nil
}

func (r *patientRepo) FindAll() ([]model.Patient, error) {
	var patients []model.Patient
	tx := r.db.Gorm.Preload("Allergies").Find(&patients)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return patients, nil
}

func (r *patientRepo) FindByID(id uint) (*model.Patient, error) {
	var patient model.Patient
	tx := r.db.Gorm.Preload("Allergies").First(&patient, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &patient, nil
}

func (r *patientRepo) Update(patient *model.Patient) (*model.Patient, error) {
	tx := r.db.Gorm.Save(patient)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = r.db.Gorm.Preload("Allergies").First(&patient, patient.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return patient, nil
}

func (r *patientRepo) Delete(id uint) error {
	var patient model.Patient
	tx := r.db.Gorm.Preload("Allergies").First(&patient, id)
	if tx.Error != nil {
		return tx.Error
	}
	err := r.db.Gorm.Model(&patient).Association("Allergies").Clear()
	if err != nil {
		return err
	}
	tx = r.db.Gorm.Delete(&model.Patient{}, id)
	return tx.Error
}

func (r *patientRepo) FilterByMRN(mrn string) (*model.Patient, error) {
	var patient model.Patient
	tx := r.db.Gorm.Preload("Allergies").Where("medical_record_number = ?", mrn).First(&patient)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &patient, nil
}

func (r *patientRepo) FindAllWithPaginationAndKeyword(
	limit int, offset int, keyword string) ([]model.Patient, int64, error) {
	var total int64
	var patients []model.Patient
	query := r.db.Gorm.Model(&model.Patient{})

	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(medical_record_number) LIKE ?",
			strings.ToLower(like), strings.ToLower(like))
	}

	tx := query.Count(&total)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	tx = query.Limit(limit).Offset(offset).Preload("Allergies").Find(&patients)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	return patients, total, nil
}

func (r *patientRepo) ReplaceAllergies(patient *model.Patient, allergyIDs []uint) error {
	var allergies []model.Allergy
	if len(allergyIDs) == 0 {
		return r.db.Gorm.Model(&patient).Association("Allergies").Clear()
	}
	tx := r.db.Gorm.Where("id IN ?", allergyIDs).Find(&allergies)
	if tx.Error != nil {
		return tx.Error
	}
	return r.db.Gorm.Model(&patient).Association("Allergies").Replace(allergies)
}
