package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//LoadTemplates insere os templates na váriavel templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

//ExecTemplate carrega uma página html na tela
func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}