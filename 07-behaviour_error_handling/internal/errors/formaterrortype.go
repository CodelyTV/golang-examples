package errors

import (
	"github.com/pkg/errors"
)

type formatDataError struct {
	error
}

func WrapFormatDataError(err error, format string, args ...interface{}) error {
	return &formatDataError{errors.Wrapf(err, format, args...)}
}

func NewFormatDataError(format string, args ...interface{}) error {
	return &formatDataError{errors.Errorf(format, args...)}
}

func IsFormatDataError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*formatDataError)
	return ok
}
