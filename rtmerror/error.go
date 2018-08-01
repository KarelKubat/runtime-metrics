package rtmerror

import (
	"fmt"
)

// Error is the error type in runtime-metrics.
type Error struct {
	desc      string
	retryable bool
	errors    []error
}

// NewError initializes an Error return struct with a description,
// and marks the error as non-retryable. Example:
//
//  err := rtmerror.NewError("%d is the answer", 42)
func NewError(f string, a ...interface{}) *Error {
	return &Error{
		desc:      fmt.Sprintf(f, a...),
		retryable: false,
		errors:    []error{},
	}
}

// WithError adds an underlying error to the Error structure. Example:
//
//  var err1, err2 error
//  err := rtmerror.NewError("something bad happened").
//    WithError(err1).WithError(err2)
func (e *Error) WithError(err error) *Error {
	e.errors = append(e.errors, err)
	return e
}

// WithRetryable marks the error as retryable.
//
//  err := rtmerror.NewError("network error").WithRetryable(true)
func (e *Error) WithRetryable(r bool) *Error {
	e.retryable = r
	return e
}

// Retryable returns true when the error is retryable according to its
// argument.
func (e *Error) Retryable() bool {
	return e.retryable
}

// Error implements the error interface and returns the error's description,
// with underlying errors concatenated and between (). Example:
//
//  err1 := fmt.Errorf("hello %s", world)
//  err2 := fmt.Errorf("pi is not %.2f", 3.14)
//  err := rtmerror.NewError("e is not %.2f", 2.71").
//    WithError(err1).WithError(err2)
//  fmt.Printf("%v\n", err)
//  // output:
//  // e is not 2.71 (hello world) (pi is not 3.14)
func (e *Error) Error() string {
	ret := e.desc
	for _, err := range e.errors {
		ret += fmt.Sprintf(" (%v)", err)
	}
	return ret
}
