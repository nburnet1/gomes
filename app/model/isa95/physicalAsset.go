package isa95

import (
	"time"

	"gorm.io/gorm"
)

type PhysicalAssetClass struct {
	gorm.Model
	Manufacturer string
	ModelID      string
	Description  string
}

type PhysicalAsset struct {
	gorm.Model
	Description          string
	Location             string
	FixedAssetID         uint
	VendorID             string
	PhysicalAssetClassID uint
	PhysicalAssetClass   PhysicalAssetClass
}

type PhysicalAssetProp struct {
	gorm.Model
	Description     string
	Value           string
	MeasurementID   uint
	PhysicalAssetID uint
	PhysicalAsset   PhysicalAsset
}

type PhysicalAssetClassProp struct {
	gorm.Model
	Description          string
	Value                string
	MeasurementID        uint
	PhysicalAssetClassID uint
	PhysicalAssetClass   PhysicalAssetClass
}

type PhysicalAssetCapabilityTest struct {
	gorm.Model
	Description string
	Version     string
}

type PhysicalAssetCapabilityTestResult struct {
	gorm.Model
	Description                   string
	Date                          time.Time
	Result                        string
	MeasurementID                 uint
	Measurement                   Measurement
	Expiration                    time.Time
	PhysicalAssetCapabilityTestID uint
	PhysicalAssetCapabilityTest   PhysicalAssetCapabilityTest
}

type EquipmentAssetLink struct {
	gorm.Model
	Description     string
	StartTime       time.Time
	EndTime         time.Time
	EquipmentID     uint
	Equipment       Equipment
	PhysicalAssetID uint
	PhysicalAsset   PhysicalAsset
}
