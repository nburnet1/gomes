package isa95

import (
	"time"
	"gomes/pkg/model"
)

type MaterialClass struct {
	model.Model
	Description            string
	AssemblyID             uint
	AssemblyRelationshipID uint
}

type MaterialClassProp struct {
	model.Model
	MaterialClassID uint
	Description     string
	Value           string
	MeasurementID   uint
}

type MaterialDefinition struct {
	model.Model
	Code                   string
	Description            string
	AssemblyTypeID         uint
	AssemblyRelationshipID uint
}

type MaterialDefinitionProp struct {
	model.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type MaterialLot struct {
	model.Model
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
	model.Model
	Code          string
	Description   string
	Value         string
	MeasurementID uint
}

type MaterialSublot struct {
	model.Model
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
	model.Model
	Code        string
	Description string
	Version     string
}

type MaterialTestResult struct {
	model.Model
	Code          string
	Description   string
	Date          time.Time
	Result        string
	MeasurementID uint
	Expiration    time.Time
}
