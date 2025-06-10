package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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
	FilterByMRN(mrn string) (*model.Patient, error)
	FindAllWithPaginationAndKeyword(limit int, offset int, keyword string) ([]model.Patient, int64, error)
	FindFromSIMRS(mrn string) (*string, *time.Time, error)
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
		Name:                strings.ToLower(request.Name),
		DateOfBirth:         request.DateOfBirth.Truncate((24 * time.Hour)),
	}

	patient, err := s.patientRepo.Create(newPatient)
	if err != nil {
		return nil, err
	}
	err = s.patientRepo.ReplaceAllergies(patient, request.AllergyIDs)
	if err != nil {
		return nil, err
	}
	patient, err = s.patientRepo.FindByID(patient.ID)
	if err != nil {
		return nil, err
	}
	return patient, err
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
		return nil, errors.New("patient not found")
	}

	patient.Name = strings.ToLower(request.Name)
	patient.MedicalRecordNumber = request.MedicalRecordNumber
	patient.DateOfBirth = request.DateOfBirth

	patient, err = s.patientRepo.Update(patient)
	if err != nil {
		return nil, err
	}
	err = s.patientRepo.ReplaceAllergies(patient, request.AllergyIDs)
	if err != nil {
		return nil, err
	}
	patient, err = s.patientRepo.FindByID(patient.ID)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (s *patientService) Delete(id uint) error {
	return s.patientRepo.Delete(id)
}

func (s *patientService) FilterByMRN(mrn string) (*model.Patient, error) {
	return s.patientRepo.FilterByMRN(mrn)
}

func (s *patientService) FindAllWithPaginationAndKeyword(
	limit int, offset int, keyword string) ([]model.Patient, int64, error) {
	return s.patientRepo.FindAllWithPaginationAndKeyword(limit, offset, keyword)
}

func (s *patientService) FindFromSIMRS(mrn string) (*string, *time.Time, error) {
	type Result struct {
		Metadata struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"metadata"`
		Response struct {
			Nomr        string      `json:"nomr"`
			NoKTP       string      `json:"no_ktp"`
			Nama        string      `json:"nama"`
			TempatLahir string      `json:"tempat_lahir"`
			TglLahir    string      `json:"tgl_lahir"`
			JnsKelamin  string      `json:"jns_kelamin"`
			NoHP        string      `json:"no_hp"`
			Kunjungan   interface{} `json:"kunjungan"`
		} `json:"response"`
	}

	req, err := http.NewRequest("GET",
		fmt.Sprintf("http://192.168.2.20/wssimrs/index.php/simrs/pasien/kunjungan/%s", mrn), nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("X-Username", os.Getenv("SIMRS_USERNAME"))
	req.Header.Set("X-Password", os.Getenv("SIMRS_PASSWORD"))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	var result Result
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, nil, err
	}

	parsedDob, err := time.Parse("2006-01-02", result.Response.TglLahir)

	return &result.Response.Nama, &parsedDob, nil

}
