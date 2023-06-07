package httperrors

import "errors"

var (
	// ErrInsertionFailed occurs during an insertion operation
	ErrInsertionFailed = errors.New("insertion operation failed")
	// ErrReadFailed occurs during a read operation
	ErrReadFailed = errors.New("read operation failed")
	// ErrUpdateFailed occurs during update operation
	ErrUpdateFailed = errors.New("update operation failed")
	// ErrDelFailed occurs during delte operation
	ErrDelFailed = errors.New("delete operation failed")
)
