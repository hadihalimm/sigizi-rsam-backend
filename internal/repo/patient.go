package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"gorm.io/gorm"
)

type PatientRepo interface {
	Create(patient *model.Patient) (*model.Patient, error)
	FindAll() ([]model.Patient, error)
	FindByID(id uint) (*model.Patient, error)
	Update(patient *model.Patient) (*model.Patient, error)
	Delete(id uint) error
	FilterByMRN(mrn string) (*model.Patient, error)
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
	tx := r.db.Gorm.Find(&patients)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return patients, nil
}

func (r *patientRepo) FindByID(id uint) (*model.Patient, error) {
	var patient model.Patient
	tx := r.db.Gorm.First(&patient, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &patient, nil
}

func (r *patientRepo) Update(patient *model.Patient) (*model.Patient, error) {
	tx := r.db.Gorm.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&patient)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return patient, nil
}

func (r *patientRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.Patient{}, id)
	return tx.Error
}

func (r *patientRepo) FilterByMRN(mrn string) (*model.Patient, error) {
	var patient model.Patient
	tx := r.db.Gorm.Where("medical_record_number = ?", mrn).First(&patient)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &patient, nil
}
