package isa95

import (
	"gorm.io/gorm"
)

type OperationType struct {
	gorm.Model
	Code        string
	Description string
}

type OperationMaterialBillItem struct {
	gorm.Model
	Code                   string
	Description            string
	MaterialClassID        uint
	MaterialDefinitionID   uint
	UseID                  uint
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
	Quantity               string
	MeasurementID          uint
}
