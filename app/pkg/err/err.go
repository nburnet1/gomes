package err

import (
	"go.uber.org/zap"
)

type Error struct {
	err error
	zap *zap.Logger
}

func NewError(zap *zap.Logger) *Error {
	return &Error{
		err: nil,
		zap: zap,
	}
}

func (e *Error) ErrIsNil(err error) bool {
	return err == nil
}


func (e *Error) SetErrorAndFatalIfNotNil(err error) {
	e.err = err
	if !e.ErrIsNil(err) {
		e.zap.Fatal(err.Error())
	}
}


