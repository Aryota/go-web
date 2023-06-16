package main

import (
	"html/template"
)

func parseTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}
