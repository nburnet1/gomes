package operation

import (
	"gomes/model/isa95"

)

type OperationProp struct {
	Code               string
	Description        string
	Value              string
	ValueMeasurementID uint
	ValueMeasurement   isa95.Measurement `gorm:"foreignKey:ValueMeasurementID"`
	Quantity           string
	QuantityMeasurementID uint
	QualityMeasurement isa95.Measurement `gorm:"foreignKey:QualityMeasurementID"`
}