package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(write http.ResponseWriter, request *http.Request){
	renderTemplate(write, "home.page.tmpl")
}

// About is the about page handler
func About(write http.ResponseWriter, request *http.Request){
	
}

func renderTemplate(write http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(write, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}


func main() {
	
	// basic routing, for / call function Home
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application server on port %s", portNumber))
	
	_ = http.ListenAndServe(portNumber, nil)
}