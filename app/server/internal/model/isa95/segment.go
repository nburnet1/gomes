package isa95

import (
	"time"
	"gomes/server/internal/model"
)

type Utilize struct {
	model.Model
	Code          string
	Description   string
	UseCategoryID uint
}

type UtilizeCategory struct {
	model.Model
	Code        string
	Description string
}

type ProcessSegment struct {
	model.Model
	Code            uint
	Description     string
	OperationTypeID uint
	LevelID         uint
	Duration        string
	MeasurementID   uint
}

type ProcessSegmentParameter struct {
	model.Model
	Description   string
	Value         string
	MeasurementID uint
}

type ProcessSegmentDependency struct {
	model.Model
	Code             string
	Description      string
	DependencyType   string
	DependencyFactor string
	MeasurementID    uint
}

type ProcessSegmentCapability struct {
	model.Model
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
	model.Model
	PersonnelID   uint
	Description   string
	Quantity      string
	MeasurementID uint
	UseID         uint
}

type PersonnelSegmentProp struct {
	model.Model
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}

type EquipmentSegment struct {
	model.Model
	EquipmentClassID uint
	Description      string
	Quantity         string
	MeasurementID    uint
	UseID            uint
}

type EquipmentSegmentProp struct {
	model.Model
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}

type MaterialSegment struct {
	model.Model
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
	model.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	UseID                uint
	Quantity             string
	MeasurementID        uint
}

type PhysicalAssetSegmentProp struct {
	model.Model
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}
