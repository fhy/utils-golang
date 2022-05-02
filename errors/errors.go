package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	NoType = ErrorType(iota)
	BadRequest
	NotFound
	//add any type you want
)

type ErrorType uint

type customError struct {
	errorType     ErrorType
	originalError error
	contextInfo   errorContext
}

// Error returns the mssage of a customError
func (error customError) Error() string {
	return error.originalError.Error()
}

// New creates a new customError
func (typ ErrorType) New(msg string) error {
	return customError{errorType: typ, originalError: errors.New(msg)}
}

// New creates a new customError with formatted message
func (typ ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)

	return customError{errorType: typ, originalError: err}
}

// Wrap creates a new wrapped error
func (typ ErrorType) Wrap(err error, msg string) error {
	return typ.Wrapf(err, msg)
}

// Wrap creates a new wrapped error with formatted message
func (typ ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)

	return customError{errorType: typ, originalError: newErr}
}

// New creates a no type error
func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) error {
	return customError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrap wrans an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf wraps an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			contextInfo:   customErr.contextInfo,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError}
}

type errorContext struct {
	Field   string
	Message string
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(customError); ok {
		return customError{errorType: customErr.errorType, originalError: customErr.originalError, contextInfo: context}
	}

	return customError{errorType: NoType, originalError: err, contextInfo: context}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customErr, ok := err.(customError); ok || customErr.contextInfo != emptyContext {

		return map[string]string{"field": customErr.contextInfo.Field, "message": customErr.contextInfo.Message}
	}

	return nil
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}
