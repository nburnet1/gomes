package isa95

import (
	"gomes/server/internal/model"
)

type Operation struct {
	model.Model
	Code             string
	Description      string
	Version          string
	OperationTypeID  uint
	LevelID          uint
	BillOfMaterialID string
	WorkDefinitionID string
	BillOfResourceID string
}

type OperationMaterialBill struct {
	model.Model
	Code        string
	Description string
}

type OperationSegment struct {
	model.Model
	Code             string
	Description      string
	LevelID          uint
	Duration         string
	MeasurementID    uint
	Measurement      Measurement
	ProcessSegmentID uint
	OperationTypeID  uint
	WorkDefinitionID uint
}

type OperationSegmentDependency struct {
	model.Model
	Code             string
	Description      string
	DependencyType   string
	DependencyFactor string
	MeasurementID    uint
	Measurement      Measurement
}

type ParameterSpecification struct {
	model.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type PersonnelSpecification struct {
	model.Model
	PersonnelClassID uint
	PersonID         uint
	Description      string
	UseID            uint
	Quantity         string
	MeasurementID    uint
	Measurement      Measurement
}

type PersonnelSpecificationProp struct {
	model.Model
	// OperationProp
}

type EquipmentSpecification struct {
	model.Model
	EquipmentClassID uint
	EquipmentID      uint
	Description      string
	UseID            uint
	Quantity         string
	MeasurementID    uint
	Measurement      Measurement
}

type EquipmentSpecificationProp struct {
	model.Model
	// OperationProp
}

type PhysicalAssetSpecification struct {
	model.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	UseID                uint
	Quantity             string
	MeasurementID        uint
}

type PhysicalAssetSpecificationProp struct {
	model.Model
	// OperationProp
}

type MaterialSpecification struct {
	model.Model
	MaterialClassID        uint
	MaterialDefinitionID   uint
	Description            string
	UseID                  uint
	Quantity               string
	MeasurementID          uint
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
}

type MaterialSpecificationProp struct {
	model.Model
	// OperationProp
}
