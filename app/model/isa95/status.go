package isa95

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Name             string
	Description      string
	StatusCategoryID uint
}

type StatusCategory struct {
	gorm.Model
	Name        string
	Description string
}
