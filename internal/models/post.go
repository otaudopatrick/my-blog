package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255)"`
	Slug        string `gorm:"unique;type:varchar(255)"`
	Body        string `gorm:"type:text"`
	Description string `gorm:"type:text"`
	Author      string `gorm:"type:varchar(255)"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
