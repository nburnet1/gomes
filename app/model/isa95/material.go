package isa95

import (
	"time"

	"gorm.io/gorm"
)

type MaterialClass struct {
	gorm.Model
	Description            string
	AssemblyID             uint
	AssemblyRelationshipID uint
}

type MaterialClassProp struct {
	gorm.Model
	MaterialClassID uint
	Description     string
	Value           string
	MeasurementID   uint
}

type MaterialDefinition struct {
	gorm.Model
	Code                   string
	Description            string
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
}

type MaterialDefinitionProp struct {
	gorm.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type MaterialLot struct {
	gorm.Model
	Code                   string
	Description            string
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
	Quantity               string
	MeasurementID          uint
	LevelID                uint
	StatusID               uint
}

type MaterialLotProp struct {
	gorm.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type MaterialSublot struct {
	gorm.Model
	Code                   string
	Description            string
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
	Quantity               string
	MeasurementID          uint
	LevelID                uint
	StatusID               uint
}

type MaterialTest struct {
	gorm.Model
	Code        string
	Description string
	Version     string
}

type MaterialTestResult struct {
	gorm.Model
	Code          string
	Description   string
	Date          time.Time
	Result        string
	MeasurementID uint
	Expiration    time.Time
}
