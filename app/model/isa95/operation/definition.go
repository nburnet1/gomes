package operation

import (
	"gomes/model/isa95"

	"gorm.io/gorm"
)

type Operation struct {
	gorm.Model
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
	gorm.Model
	Code        string
	Description string
}

type OperationSegment struct {
	gorm.Model
	Code             string
	Description      string
	LevelID          uint
	Duration         string
	MeasurementID    uint
	Measurement      isa95.Measurement
	ProcessSegmentID uint
	OperationTypeID  uint
	WorkDefinitionID uint
}

type OperationSegmentDependency struct {
	gorm.Model
	Code             string
	Description      string
	DependencyType   string
	DependencyFactor string
	MeasurementID    uint
	Measurement      isa95.Measurement
}

type ParameterSpecification struct {
	gorm.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type PersonnelSpecification struct {
	gorm.Model
	PersonnelClassID uint
	PersonID         uint
	Description      string
	UseID            uint
	Quantity         string
	MeasurementID    uint
	Measurement      isa95.Measurement
}

type PersonnelSpecificationProp struct {
	gorm.Model
	OperationProp
}

type EquipmentSpecification struct {
	gorm.Model
	EquipmentClassID uint
	EquipmentID      uint
	Description      string
	UseID            uint
	Quantity         string
	MeasurementID    uint
	Measurement      isa95.Measurement
}

type EquipmentSpecificationProp struct {
	gorm.Model
	OperationProp
}

type PhysicalAssetSpecification struct {
	gorm.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	UseID                uint
	Quantity             string
	MeasurementID        uint
}

type PhysicalAssetSpecificationProp struct {
	gorm.Model
	OperationProp
}

type MaterialSpecification struct {
	gorm.Model
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
	gorm.Model
	OperationProp
}
