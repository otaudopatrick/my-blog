package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
