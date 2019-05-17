package errors

import (
	"github.com/pkg/errors"
)

type dataUnreacheable struct {
	error
}

// WrapDataUnreacheable returns an error which wraps err that satisfies
// IsDataUnreacheable()
func WrapDataUnreacheable(err error, format string, args ...interface{}) error {
	return &dataUnreacheable{errors.Wrapf(err, format, args...)}
}

// NewDataUnreacheable returns an error which satisfies IsDataUnreacheable()
func NewDataUnreacheable(format string, args ...interface{}) error {
	return &dataUnreacheable{errors.Errorf(format, args...)}
}

// IsDataUnreacheable reports whether err was created with DataUnreacheablef() or
// NewDataUnreacheable()
func IsDataUnreacheable(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*dataUnreacheable)
	return ok
}
