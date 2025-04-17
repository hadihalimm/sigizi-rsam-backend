package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type PatientService interface {
	Create(request request.CreatePatient) (*model.Patient, error)
	GetAll() ([]model.Patient, error)
	GetByID(id uint) (*model.Patient, error)
	Update(id uint, request request.UpdatePatient) (*model.Patient, error)
	Delete(id uint) error
}

type patientService struct {
	patientRepo repo.PatientRepo
	validate    *validator.Validate
}

func NewPatientService(patientRepo repo.PatientRepo, validate *validator.Validate) PatientService {
	return &patientService{patientRepo: patientRepo, validate: validate}
}

func (s *patientService) Create(request request.CreatePatient) (*model.Patient, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newPatient := &model.Patient{
		MedicalRecordNumber: request.MedicalRecordNumber,
		Name:                request.Name,
		DateOfBirth:         request.DateOfBirth,
	}
	return s.patientRepo.Create(newPatient)
}

func (s *patientService) GetAll() ([]model.Patient, error) {
	return s.patientRepo.FindAll()
}

func (s *patientService) GetByID(id uint) (*model.Patient, error) {
	return s.patientRepo.FindByID(id)
}

func (s *patientService) Update(id uint, request request.UpdatePatient) (*model.Patient, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	patient, err := s.patientRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	patient.Name = request.Name
	patient.MedicalRecordNumber = request.MedicalRecordNumber
	patient.DateOfBirth = request.DateOfBirth
	return s.patientRepo.Update(patient)
}

func (s *patientService) Delete(id uint) error {
	return s.patientRepo.Delete(id)
}
