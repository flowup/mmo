package utils

import (
	"io/ioutil"
	"path"
)

// GetTemplateHelp is function to extract help (list of options, usage) from the template
func GetTemplateHelp(template string) string {
	b, _ := ioutil.ReadFile(path.Join(template, "__help.txt"))
	return string(b)
}
