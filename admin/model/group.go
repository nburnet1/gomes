package model

import "gomes/namespace"

func init(){
	namespace.RegisterModels(Group{})
}

type Group struct {
	GroupID int `gorm:"primaryKey"`
	Name    string
}