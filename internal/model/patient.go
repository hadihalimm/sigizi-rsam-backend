package model

import "time"

type Patient struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	MedicalRecordNumber string    `gorm:"unique;not null" json:"medicalRecordNumber"`
	Name                string    `gorm:"not null" json:"name"`
	DateOfBirth         time.Time `gorm:"not null" json:"dateOfBirth"`
}
