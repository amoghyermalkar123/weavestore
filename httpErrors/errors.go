package httperrors

import "errors"

var (
	ErrInsertionFailed = errors.New("insertion operation failed")
	ErrReadFailed      = errors.New("read operation failed")
	ErrUpdateFailed    = errors.New("update operation failed")
	ErrDelFailed       = errors.New("delete operation failed")
)
