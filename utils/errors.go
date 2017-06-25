package utils

import "errors"

var (
	ErrNotImplemented = errors.New("Not implemented yet")

	ErrNoArg = errors.New("No arguments")

	ErrParseFileName = errors.New("Can't parse file name from path")
)
