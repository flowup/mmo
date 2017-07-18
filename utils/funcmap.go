package utils

import (
	"text/template"
	"strings"
)

var (
	//DefaultFuncMap is a map of default functions available within templates
	DefaultFuncMap = template.FuncMap{
		"Title": strings.Title,
	}
)
