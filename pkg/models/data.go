package models

import (
	"gorm.io/gorm"
	"time"
)

type Data struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Field1    string `json:"field1,omitempty"`
	Field2    string `json:"field2,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
