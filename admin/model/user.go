package model

import "gomes/namespace"


func init(){
	namespace.RegisterModels(User{})
}


type User struct {
	UserID   int `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
	GroupID  int
	Group   Group
}