package models

import (
	"gorm.io/gorm"
	"time"
)

type Vehicle struct {
	gorm.Model
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Make      string `json:"make"`
	ModelType string `json:"modelType"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
