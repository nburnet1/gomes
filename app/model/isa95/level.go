package isa95

import (
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model
	Description   string
	LevelLookupID uint
	LevelLookup   LevelLookup `gorm:"foreignKey:LevelLookupID"`
	ParentLevelID uint
}

type LevelLookup struct {
	gorm.Model
	Code uint
	Name string
}
