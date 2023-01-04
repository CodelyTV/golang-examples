package errors

import (
	"github.com/pkg/errors"
)

type httpDataError struct {
	error
}

func WrapHttpDataError(err error, format string, args ...interface{}) error {
	return &httpDataError{errors.Wrapf(err, format, args...)}
}

func NewHttpDataError(format string, args ...interface{}) error {
	return &httpDataError{errors.Errorf(format, args...)}
}

func IsHttpDataError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*httpDataError)
	return ok
}
