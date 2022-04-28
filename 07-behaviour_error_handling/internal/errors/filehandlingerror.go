package errors

import (
	"github.com/pkg/errors"
)

type fileHandlingError struct {
	error
}

// WrapFileHandlingError returns an error which wraps err that satisfies
// IsFileHandlingError()
func WrapFileHandlingError(err error, format string, args ...interface{}) error {
	return &fileHandlingError{errors.Wrapf(err, format, args...)}
}

// NewFileHandlingError returns an error which satisfies IsFileHandlingError()
func NewFileHandlingError(format string, args ...interface{}) error {
	return &fileHandlingError{errors.Errorf(format, args...)}
}

// IsFileHandlingError reports whether err was created with FileHandlingError() or
// NewFileHandlingError()
func IsFileHandlingError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*fileHandlingError)
	return ok
}
