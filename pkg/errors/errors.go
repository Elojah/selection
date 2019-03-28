package errors

import (
	"fmt"
)

// #HTTP errors

// ErrMissingParam is raised when a mandatory param is missing.
type ErrMissingParam struct {
	Name string
}

// Error cast for ErrMissingParam.
func (e ErrMissingParam) Error() string {
	return fmt.Sprintf("missing parameter %s", e.Name)
}

// #Dev errors

// ErrNotImplemented is raised when a non implemented feature is called.
type ErrNotImplemented struct {
	Version string
}

// Error cast for ErrNotImplemented.
func (e ErrNotImplemented) Error() string {
	return fmt.Sprintf("not implemented before version %s", e.Version)
}
