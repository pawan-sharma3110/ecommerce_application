package main

import (
	"embed"
	"html/template"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	Api             string
	cssVersion      string
}

var function = template.FuncMap{}


var templateFS embed.FS

