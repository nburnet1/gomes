package isa95

import (
	"gomes/pkg/model"
)

type OperationType struct {
	model.Model
	Code        string
	Description string
}

type OperationMaterialBillItem struct {
	model.Model
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
