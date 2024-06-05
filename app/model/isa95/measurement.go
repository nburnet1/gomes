package isa95

import (
	"gorm.io/gorm"
)

type Measurement struct {
	gorm.Model
	Code string
	Name string
	Description string
}