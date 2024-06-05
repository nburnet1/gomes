package operation

import (
	"gomes/model/isa95"
	"time"

	"gorm.io/gorm"
)

type OperationPerformance struct {
	gorm.Model
	Code                string
	Description         string
	OperationTypeID     uint
	OperationScheduleID uint
	StartTime           time.Time
	EndTime             time.Time
	LevelID             uint
	PerformanceState    string
}

type OperationResponse struct {
	gorm.Model
	Code                  string
	Description           string
	OperationTypeID       uint
	OperationRequestID    uint
	StartTime             time.Time
	EndTime               time.Time
	LevelID               uint
	OperationDefinitionID uint
	ResponseState         string
}

type SegmentResponse struct {
	gorm.Model
	Code                  string
	Description           string
	OperationTypeID       uint
	ProcessSegmentID      uint
	ActualStartTime       time.Time
	ActualEndTime         time.Time
	LevelID               uint
	OperationDefinitionID uint
}

type SegmentData struct {
	gorm.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type PersonnelActual struct {
	gorm.Model
	PersonnelClassID uint
	PersonID         uint
	Description      string
	UseID            uint
	Quantity         string
	MeasurementID    uint
	Measurement      isa95.Measurement
}

type PersonnelActualProp struct {
	gorm.Model
	OperationProp
}

type EquipmentActual struct {
	gorm.Model
	EquipmentClassID uint
	EquipmentID      uint
	Description      string
	UseID            uint
	Quantity         uint
	MeasurementID    uint
	Measurement      isa95.Measurement
}

type EquipmentActualProperty struct {
	gorm.Model
	OperationProp
}

type PhysicalAssetActual struct {
	gorm.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	UseID                uint
	Quantity             uint
	MeasurementID        uint
}

type PhysicalAssetActualProperty struct {
	gorm.Model
	OperationProp
}

type MaterialActual struct {
	gorm.Model
	MaterialClassID        uint
	MaterialDefinitionID   uint
	MaterialLotID          uint
	MaterialSublotID       uint
	Description            string
	UseID                  uint
	LevelID                uint
	Quantity               string
	MeasurementID          uint
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
}

type MaterialActualProperty struct {
	gorm.Model
	OperationProp
}
