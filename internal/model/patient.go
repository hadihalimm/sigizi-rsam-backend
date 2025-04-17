package model

import "time"

type Patient struct {
	ID                  uint      `gorm:"primaryKey"`
	MedicalRecordNumber string    `gorm:"unique;not null"`
	Name                string    `gorm:"not null"`
	DateOfBirth         time.Time `gorm:"not null"`
}
