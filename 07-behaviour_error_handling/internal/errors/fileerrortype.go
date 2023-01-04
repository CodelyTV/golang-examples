package errors

import (
	"github.com/pkg/errors"
)

type fileDataError struct {
	error
}

func WrapFileDataUError(err error, format string, args ...interface{}) error {
	return &fileDataError{errors.Wrapf(err, format, args...)}
}

func NewFileDataError(format string, args ...interface{}) error {
	return &fileDataError{errors.Errorf(format, args...)}
}

func IsFileErrorType(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*fileDataError)
	return ok
}
