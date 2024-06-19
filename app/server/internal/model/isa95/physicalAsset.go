package isa95

import (
	"time"
	"gomes/server/internal/model"
)

type PhysicalAssetClass struct {
	model.Model
	Manufacturer string
	ModelID      string
	Description  string
}

type PhysicalAsset struct {
	model.Model
	Description          string
	Location             string
	FixedAssetID         uint
	VendorID             string
	PhysicalAssetClassID uint
	PhysicalAssetClass   PhysicalAssetClass
}

type PhysicalAssetProp struct {
	model.Model
	Description     string
	Value           string
	MeasurementID   uint
	PhysicalAssetID uint
	PhysicalAsset   PhysicalAsset
}

type PhysicalAssetClassProp struct {
	model.Model
	Description          string
	Value                string
	MeasurementID        uint
	PhysicalAssetClassID uint
	PhysicalAssetClass   PhysicalAssetClass
}

type PhysicalAssetCapabilityTest struct {
	model.Model
	Description string
	Version     string
}

type PhysicalAssetCapabilityTestResult struct {
	model.Model
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
	model.Model
	Description     string
	StartTime       time.Time
	EndTime         time.Time
	EquipmentID     uint
	Equipment       Equipment
	PhysicalAssetID uint
	PhysicalAsset   PhysicalAsset
}
