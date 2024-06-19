package isa95

import (
	"gomes/server/internal/model"

	"time"
)

type OperationPerformance struct {
	model.Model
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
	model.Model
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
	model.Model
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
	model.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type PersonnelActual struct {
	model.Model
	PersonnelClassID uint
	PersonID         uint
	Description      string
	UseID            uint
	Quantity         string
	MeasurementID    uint
	Measurement      Measurement
}

type PersonnelActualProp struct {
	model.Model
	// OperationProp
}

type EquipmentActual struct {
	model.Model
	EquipmentClassID uint
	EquipmentID      uint
	Description      string
	UseID            uint
	Quantity         uint
	MeasurementID    uint
	Measurement      Measurement
}

type EquipmentActualProperty struct {
	model.Model
	// OperationProp
}

type PhysicalAssetActual struct {
	model.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	UseID                uint
	Quantity             uint
	MeasurementID        uint
}

type PhysicalAssetActualProperty struct {
	model.Model
	// OperationProp
}

type MaterialActual struct {
	model.Model
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
	model.Model
	// OperationProp
}
