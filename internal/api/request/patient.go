package request

import "time"

type CreatePatient struct {
	MedicalRecordNumber string    `json:"medicalRecordNumber" binding:"required"`
	Name                string    `json:"name" binding:"required"`
	DateOfBirth         time.Time `json:"dateOfBirth" binding:"required"`
	AllergyIDs          []uint    `json:"allergyIDs"`
}

type UpdatePatient struct {
	MedicalRecordNumber string    `json:"medicalRecordNumber" binding:"required"`
	Name                string    `json:"name" binding:"required"`
	DateOfBirth         time.Time `json:"dateOfBirth" binding:"required"`
	AllergyIDs          []uint    `json:"allergyIDs"`
}
