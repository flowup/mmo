package utils

import "errors"

// Errors returned by the mmo tool
var (
	ErrNotImplemented = errors.New("Not implemented yet")

	ErrNoArg = errors.New("No arguments")

	ErrNoProject = errors.New("No project initialized in current directory")

	ErrServiceNotExists = errors.New("Following service doesn't exist: ")

	ErrContextNotSet = errors.New("Context not set. Please set context using sub-command \"set-context\"")
)
