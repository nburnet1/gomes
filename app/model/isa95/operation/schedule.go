package operation

import (
	"gomes/model/isa95"
	"time"

	"gorm.io/gorm"
)

type OperationSchedule struct {
	gorm.Model
	Code            string
	Description     string
	OperationTypeID uint
	StartTime       time.Time
	EndTime         time.Time
	LevelID         uint
	ScheduleState   string
}

type OperationRequest struct {
	gorm.Model
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
	gorm.Model
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
	gorm.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type PersonnelRequirement struct {
	gorm.Model
	PersonnelClassID uint
	PersonID         uint
	Description      string
	UseID            uint
	Quantity         string
	MeasurementID    uint
	Measurement      isa95.Measurement
}

type PersonnelRequirementProp struct {
	gorm.Model
	OperationProp
}

type PhysicalAssetRequirement struct {
	gorm.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	UseID                uint
	Quantity             uint
	MeasurementID        uint
	LevelID              uint
}

type PhysicalAssetRequirementProperty struct {
	gorm.Model
	OperationProp
}

type MaterialRequirement struct {
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

type MaterialRequirementProperty struct {
	gorm.Model
	OperationProp
}
