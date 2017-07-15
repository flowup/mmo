package utils

import "os"

// CreateDir Create directory in path
func CreateDir(path string) error{
	return os.MkdirAll(path,  os.ModePerm)
}
