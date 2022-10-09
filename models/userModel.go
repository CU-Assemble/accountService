package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	StudentId string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Nickname  string
	Faculty   string
	Tel       string
	Email     string
	Password  string
}
