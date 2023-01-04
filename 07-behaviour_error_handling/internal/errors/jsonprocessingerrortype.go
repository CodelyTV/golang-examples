package errors

import (
	"github.com/pkg/errors"
)

type jsonProcessingDataError struct {
	error
}

func WrapJsonProcessingDataError(err error, format string, args ...interface{}) error {
	return &jsonProcessingDataError{errors.Wrapf(err, format, args...)}
}

func NewJsonProcessingDataError(format string, args ...interface{}) error {
	return &jsonProcessingDataError{errors.Errorf(format, args...)}
}

func IsJsonProcessingDataError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*jsonProcessingDataError)
	return ok
}
