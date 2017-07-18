package utils

import (
	"strings"
	"text/template"
)

var (
	//DefaultFuncMap is a map of default functions available within templates
	DefaultFuncMap = template.FuncMap{
		"Title": strings.Title,
	}
)
