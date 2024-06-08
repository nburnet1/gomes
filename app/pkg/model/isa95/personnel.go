package isa95

import (
	"time"
	"gomes/pkg/model"
)

// PersonnelClass
// @Description User account information
type PersonnelClass struct {
	model.Model
	Code        string `json:"code" example:"asdgw34"`
	Description string `json:"description" example:"Engineering Team"`
}

type PersonnelClassProp struct {
	model.Model
	PersonnelClassID uint
	PersonnelClass   PersonnelClass
	Description      string
	Value            string
	MeasurementID    uint
}

type Person struct {
	model.Model
	Description      string
	Name             string
	PersonnelClassID uint
	PersonnelClass   PersonnelClass
}

type PersonProp struct {
	model.Model
	PersonID      uint
	Person        Person
	Description   string
	Value         string
	MeasurementID uint
}

type QualificationTest struct {
	model.Model
	Description string
	Version     string
}

type QualificationTestProp struct {
	model.Model
	Description         string
	Date                time.Time
	Result              bool
	QualificationTestID uint
	QualificationTest   QualificationTest
	Expiration          time.Time
}
