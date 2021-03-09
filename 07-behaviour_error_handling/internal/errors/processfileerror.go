package errors

import (
	"github.com/pkg/errors"
)

type processFileError struct {
	error
}

// WrapProcessFileError returns an error which wraps err that satisfies
// IsProcessFileError()
func WrapProcessFileError(err error, format string, args ...interface{}) error {
	return &processFileError{errors.Wrapf(err, format, args...)}
}

// NewProcessFileError returns an error which satisfies IsProcessFileError()
func NewProcessFileError(format string, args ...interface{}) error {
	return &processFileError{errors.Errorf(format, args...)}
}

// IsProcessFileError reports whether err was created with WrapProcessFileError() or
// NewProcessFileError()
func IsProcessFileError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*processFileError)
	return ok
}
