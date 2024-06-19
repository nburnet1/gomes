package isa95

import (
	"gomes/server/internal/model"
	"time"
)

type OperationSchedule struct {
	model.Model
	Code            string
	Description     string
	OperationTypeID uint
	StartTime       time.Time
	EndTime         time.Time
	LevelID         uint
	ScheduleState   string
}

type OperationRequest struct {
	model.Model
	Code                  string
	Description           string
	OperationTypeID       uint
	StartTime             time.Time
	EndTime               time.Time
	Priority              uint
	LevelID               uint
	OperationDefinitionID uint
	ScheduleState         string
}

type SegmentRequirement struct {
	model.Model
	Code                  string
	Description           string
	OperationTypeID       uint
	SegmentID             uint
	EarliestStartTime     time.Time
	LatestEndTime         time.Time
	Duration              string
	MeasurementID         uint
	LevelID               uint
	OperationDefinitionID uint
}

type SegmentParameter struct {
	model.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type PersonnelRequirement struct {
	model.Model
	PersonnelClassID uint
	PersonID         uint
	Description      string
	UseID            uint
	Quantity         string
	MeasurementID    uint
	Measurement      Measurement
}

type PersonnelRequirementProp struct {
	model.Model
	// OperationProp
}

type PhysicalAssetRequirement struct {
	model.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	UseID                uint
	Quantity             uint
	MeasurementID        uint
	LevelID              uint
}

type PhysicalAssetRequirementProperty struct {
	model.Model
	// OperationProp
}

type MaterialRequirement struct {
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

type MaterialRequirementProperty struct {
	model.Model
	// OperationProp
}
