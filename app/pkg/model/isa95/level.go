package isa95

import (
	"gomes/pkg/model"
)

type Level struct {
	model.Model
	Description   string
	LevelLookupID uint
	LevelLookup   LevelLookup `gorm:"foreignKey:LevelLookupID"`
	ParentLevelID uint
}

type LevelLookup struct {
	model.Model
	Code uint
	Name string
}
