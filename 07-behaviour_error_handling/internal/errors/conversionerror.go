package errors

import (
	"github.com/pkg/errors"
)

type conversionError struct {
	error
}

// WrapConversionError returns an error which wraps err that satisfies
// IsConversionError()
func WrapConversionError(err error, format string, args ...interface{}) error {
	return &conversionError{errors.Wrapf(err, format, args...)}
}

// NewConversionError returns an error which satisfies IsConversionError()
func NewConversionError(format string, args ...interface{}) error {
	return &conversionError{errors.Errorf(format, args...)}
}

// IsConversionError reports whether err was created with ConversionError() or
// NewConversionError()
func IsConversionError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*conversionError)
	return ok
}
