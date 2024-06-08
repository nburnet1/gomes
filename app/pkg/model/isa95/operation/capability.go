package isa95

import (
	"gomes/pkg/model"
	"gomes/pkg/model/isa95"
	"time"
)

type OperationCapability struct {
	model.Model
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
	model.Model
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
	model.Model
	OperationProp
}

type EquipmentCapability struct {
	model.Model
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
	model.Model
	OperationProp
}

type PhysicalAssetCapability struct {
	model.Model
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
	model.Model
	OperationProp
}

type MaterialCapability struct {
	model.Model
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
	model.Model
	OperationProp
}
