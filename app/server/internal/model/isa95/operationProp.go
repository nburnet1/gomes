package isa95

import (
)

type OperationProp struct {
	Code                  string
	Description           string
	Value                 string
	ValueMeasurementID    uint
	ValueMeasurement      Measurement `gorm:"foreignKey:ValueMeasurementID"`
	Quantity              string
	QuantityMeasurementID uint
	QualityMeasurement    Measurement `gorm:"foreignKey:QualityMeasurementID"`
}
