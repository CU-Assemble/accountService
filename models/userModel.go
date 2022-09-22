package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StudentId string
	Name      string
	Nickname  string
	Birthdate time.Time
	Faculty   string
	Tel       string
	Email     string
}
