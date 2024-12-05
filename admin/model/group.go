package model

import "github.com/nburnet1/gomes/pkg/namespace"

func init(){
	namespace.RegisterModels(Group{})
}

type Group struct {
	GroupID int `gorm:"primaryKey"`
	Name    string
}