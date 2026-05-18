package models

import (
	"time"

	"gorm.io/gorm"
)

type Feedback struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);not null" json:"email"`
	Subject   string         `gorm:"type:varchar(255);not null" json:"subject"`
	Message   string         `gorm:"type:text;not null" json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Feedback) TableName() string {
	return "feedback"
}
