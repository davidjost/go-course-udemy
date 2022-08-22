package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(write http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(write, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}