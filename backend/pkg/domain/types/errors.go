package types

import "github.com/m-mizutani/goerr"

var (
	ErrRecordNotFound     = goerr.New("Record not found")
	ErrFailedPrecondition = goerr.New("Failed precondition")
	ErrInvalidArgument    = goerr.New("Invalid argument")
	ErrPermissionDenied   = goerr.New("Permission denied")
)

func ErrRecordNotFound2(msg string) *goerr.Error {
	return goerr.New(msg)
}
