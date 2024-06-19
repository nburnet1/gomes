package isa95

import "gomes/server/internal/model"

type Measurement struct {
	model.Model
	Code        string
	Name        string
	Description string
}
