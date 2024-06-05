package isa95

import (
	"gorm.io/gorm"
)

type AssemblyType struct {
	gorm.Model
	Code uint
	Name string
}

type AssemblyRelationship struct {
	gorm.Model
	Code uint
	Name string
}
