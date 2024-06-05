package isa95

import (
	"time"

	"gorm.io/gorm"
)

type Use struct {
	gorm.Model
	Code          string
	Description   string
	UseCategoryID uint
}

type UseCategory struct {
	gorm.Model
	Code        string
	Description string
}

type ProcessSegment struct {
	gorm.Model
	Code            uint
	Description     string
	OperationTypeID uint
	LevelID         uint
	Duration        string
	MeasurementID   uint
}

type ProcessSegmentParameter struct {
	gorm.Model
	Description   string
	Value         string
	MeasurementID uint
}

type ProcessSegmentDependency struct {
	gorm.Model
	Code             string
	Description      string
	DependencyType   string
	DependencyFactor string
	MeasurementID    uint
}

type ProcessSegmentCapability struct {
	gorm.Model
	Code             string
	Description      string
	ProcessSegmentID uint
	CapacityType     string
	Reason           string
	LevelID          uint
	StartTime        time.Time
	EndTime          time.Time
}

type PersonnelSegment struct {
	gorm.Model
	PersonnelID   uint
	Description   string
	Quantity      string
	MeasurementID uint
	UseID         uint
}

type PersonnelSegmentProp struct {
	gorm.Model
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}

type EquipmentSegment struct {
	gorm.Model
	EquipmentClassID uint
	Description      string
	Quantity         string
	MeasurementID    uint
	UseID            uint
}

type EquipmentSegmentProp struct {
	gorm.Model
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}

type MaterialSegment struct {
	gorm.Model
	MaterialClassID        uint
	MaterialDefinitionID   uint
	Description            string
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
	UseID                  uint
}

type MaterialSegmentProp struct {
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}

type PhysicalAssetSegment struct {
	gorm.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	UseID                uint
	Quantity             string
	MeasurementID        uint
}

type PhysicalAssetSegmentProp struct {
	gorm.Model
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}
