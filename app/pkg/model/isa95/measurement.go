package isa95

import "gomes/pkg/model"

type Measurement struct {
	model.Model
	Code        string
	Name        string
	Description string
}
