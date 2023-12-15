package stakego

import (
	"fmt"
)

// NewStakeError - create a wrapped error
func NewStakeError(msg string, err error) error {
	e := StakeError{}
	e.msg = msg
	e.cause = err
	return e
}

// StakeError - generic error type
type StakeError struct {
	msg string
	cause error
}

// Error - error compatible message
func (e StakeError) Error() string {
	if e.msg == "" {
		return fmt.Sprintf("%v", e.cause)
	}
	return fmt.Sprintf("%s: %v", e.msg, e.cause)
}

// Unwrap the internal error
func (e StakeError) Unwrap() error {
	return e.cause
}

var (
	ErrSessionTokenMissing = NewStakeError("", fmt.Errorf("session token is invalid or missing"))
	ErrInvalidAPIResponse = NewStakeError("", fmt.Errorf("invalid API response"))
)