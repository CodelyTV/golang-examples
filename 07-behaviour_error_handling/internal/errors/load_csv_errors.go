package errors

import (
	"github.com/pkg/errors"
)

type loadCSVError struct {
	error
}

// WrapLoadCSVError returns an error which wraps err that satisfies
// IsLoadCSVError()
func WrapLoadCSVError(err error, format string, args ...interface{}) error {
	return &loadCSVError{errors.Wrapf(err, format, args...)}
}

// NewLoadCSVError returns an error which satisfies IsLoadCSVError()
func NewLoadCSVError(format string, args ...interface{}) error {
	return &loadCSVError{errors.Errorf(format, args...)}
}

// IsLoadCSVError reports whether err was created with LoadCSVErrorf() or
// NewLoadCSVError()
func IsLoadCSVError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*loadCSVError)
	return ok
}
