package model

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Name         string    `gorm:"not null" json:"name"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"type:ENUM('perawat','ahli_gizi','admin');not null" json:"role"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}
