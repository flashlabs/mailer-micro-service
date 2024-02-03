package pkg

import "errors"

var (
	ErrInvalidType    = errors.New("invalid data type")
	ErrMissingParam   = errors.New("missing param")
	ErrInvalidPayload = errors.New("invalid payload")
)
