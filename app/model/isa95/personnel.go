package isa95

import (
	"time"

	"gorm.io/gorm"
)

type PersonnelClass struct {
	gorm.Model
	Code        string
	Description string
}

type PersonnelClassProp struct {
	gorm.Model
	PersonnelClassID uint
	PersonnelClass   PersonnelClass
	Description      string
	Value            string
	MeasurementID    uint
}

type Person struct {
	gorm.Model
	Description string
	Name        string
}

type PersonProp struct {
	gorm.Model
	PersonID      uint
	Person        Person
	Description   string
	Value         string
	MeasurementID uint
}

type QualificationTest struct {
	gorm.Model
	Description string
	Version     string
}

type QualificationTestProp struct {
	gorm.Model
	Description         string
	Date                time.Time
	Result              bool
	QualificationTestID uint
	QualificationTest   QualificationTest
	Expiration          time.Time
}
