package utils

import "errors"

var (
	ErrNotImplemented = errors.New("Not implemented yet")

	ErrNoArg = errors.New("No arguments")

	ErrParseFileName = errors.New("Can't parse file name from path")

	ErrServiceNotExists = errors.New("Following service doesn't exist: ")

	ErrSetContextNoArg = errors.New("You should provide at least one argument for \"set-context\" sub-command")

	ErrContextNotSet = errors.New("Context not set. Please set context using sub-command \"set-context\"")
)
