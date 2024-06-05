package operation

import (
	"gomes/model/isa95"
	"time"

	"gorm.io/gorm"
)

type OperationCapability struct {
	gorm.Model
	Code             string
	Description      string
	CapacityType     string
	Reason           string
	ConfidenceFactor string
	StartTime        time.Time
	EndTime          time.Time
	LevelID          uint
}

type PersonnelCapability struct {
	gorm.Model
	PersonnelClassID uint
	PersonID         uint
	Description      string
	CapabilityType   string
	Reason           string
	ConfidenceFactor string
	LevelID          uint
	UseID            uint
	StartTime        time.Time
	EndTime          time.Time
	Quantity         string
	MeasurementID    uint
	Measurement      isa95.Measurement
}

type PersonnelCapabilityProp struct {
	gorm.Model
	OperationProp
}

type EquipmentCapability struct {
	gorm.Model
	EquipmentClassID uint
	EquipmentID      uint
	Description      string
	CapabilityType   string
	Reason           string
	ConfidenceFactor string
	LevelID          uint
	UseID            uint
	StartTime        time.Time
	EndTime          time.Time
	Quantity         string
	MeasurementID    uint
	Measurement      isa95.Measurement
}

type EquipmentCapabilityProp struct {
	gorm.Model
	OperationProp
}

type PhysicalAssetCapability struct {
	gorm.Model
	PhysicalAssetClassID uint
	PhysicalAssetID      uint
	Description          string
	CapabilityType       string
	Reason               string
	ConfidenceFactor     string
	LevelID              uint
	UseID                uint
	StartTime            time.Time
	EndTime              time.Time
	Quantity             uint
	MeasurementID        uint
}

type PhysicalAssetCapabilityProp struct {
	gorm.Model
	OperationProp
}

type MaterialCapability struct {
	gorm.Model
	MaterialClassID        uint
	MaterialDefinitionID   uint
	MaterialLotID          uint
	MaterialSublotID       uint
	Description            string
	CapabilityType         string
	Reason                 string
	ConfidenceFactor       string
	LevelID                uint
	UseID                  uint
	StartTime              time.Time
	EndTime                time.Time
	Quantity               string
	MeasurementID          uint
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
}

type MaterialCapabilityProp struct {
	gorm.Model
	OperationProp
}
