package zutil

import "github.com/pkg/errors"

func IncSequence() TSequence {
	sequence++
	return sequence
}
// IsTimeoutError checks if the error is a timeout error
func IsTimeoutError(err error) bool {
	if err == nil {
		return false
	}

	err = errors.Cause(err)
	ne, ok := err.(timeoutError)
	return ok && ne.Timeout()
}