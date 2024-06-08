package isa95

import (
	"gomes/pkg/model"
)

type AssemblyType struct {
	model.Model
	Code uint
	Name string
}

type AssemblyRelationship struct {
	model.Model
	Code uint
	Name string
}
