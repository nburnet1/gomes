package isa95

import (
	"gomes/pkg/model"
)

type Status struct {
	model.Model
	Name             string
	Description      string
	StatusCategoryID uint
}

type StatusCategory struct {
	model.Model
	Name        string
	Description string
}
