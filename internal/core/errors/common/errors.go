package common_errors

import "errors"

var (
	ErrBadRequest = errors.New("bag request")
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)
