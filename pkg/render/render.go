package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(write http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(write, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}