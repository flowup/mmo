package utils

import (
	"io/ioutil"
	"path"
)

func GetTemplateHelp(template string) string {
	b, _ := ioutil.ReadFile(path.Join(template, "__help.txt"))
	return string(b)
}
