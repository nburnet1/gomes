package isa95

import (
	"time"
	"gomes/server/internal/model"
)

type EquipmentClass struct {
	model.Model
	Description string
	LevelID     uint
}

type EquipmentClassProp struct {
	model.Model
	EquipmentClassID uint
	EquipmentClass   EquipmentClass
	Description      string
	Value            string
	MeasurementID    uint
}

type Equipment struct {
	model.Model
	Description string
	LevelID     uint
	Level       Level
}

type EquipmentProp struct {
	model.Model
	EquipmentID   uint
	Equipment     Equipment
	Description   string
	Value         string
	MeasurementID uint
	Measurement   Measurement
}

type EquipmentCapabilityTest struct {
	model.Model
	Description string
	Version     string
	EquipmentID uint
}

type EquipmentCapabilityTestResult struct {
	model.Model
	Description   string
	Date          time.Time
	Value         string
	MeasurementID uint
	Expiration    time.Time
}

type EquipmentRequirement struct {
	model.Model
	EquipmentClassID uint
	EquipmentID      uint
	Description      string
	UseID            uint
	Quantity         uint
	MeasurementID    uint
	LevelID          uint
}

type EquipmentRequirementProperty struct {
	model.Model
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}
