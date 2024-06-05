package isa95

import (
	"time"

	"gorm.io/gorm"
)

type EquipmentClass struct {
	gorm.Model
	Description string
	LevelID     uint
}

type EquipmentClassProp struct {
	gorm.Model
	EquipmentClassID uint
	EquipmentClass   EquipmentClass
	Description      string
	Value            string
	MeasurementID    uint
}

type Equipment struct {
	gorm.Model
	Description string
	LevelID     uint
	Level       Level
}

type EquipmentProp struct {
	gorm.Model
	EquipmentID   uint
	Equipment     Equipment
	Description   string
	Value         string
	MeasurementID uint
	Measurement   Measurement
}

type EquipmentCapabilityTest struct {
	gorm.Model
	Description string
	Version     string
	EquipmentID uint
}

type EquipmentCapabilityTestResult struct {
	gorm.Model
	Description   string
	Date          time.Time
	Value         string
	MeasurementID uint
	Expiration    time.Time
}

type EquipmentRequirement struct {
	gorm.Model
	EquipmentClassID uint
	EquipmentID      uint
	Description      string
	UseID            uint
	Quantity         uint
	MeasurementID    uint
	LevelID          uint
}

type EquipmentRequirementProperty struct {
	gorm.Model
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	Quantity              string
	QuantityMeasurementID uint
}
