package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StudentId string
	Name      string
	Nickname  string
	Faculty   string
	Tel       string
	Email     string
	Password  string
}
