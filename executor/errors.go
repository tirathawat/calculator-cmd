package executor

import "errors"

var (
	ErrRequiredArguments = errors.New("requires two arguments")
	ErrInvalidArgument   = errors.New("invalid argument")
)
