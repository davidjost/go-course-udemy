package main

import (
	"net/http"
)

// Home is the home page handler
func Home(write http.ResponseWriter, request *http.Request){
	renderTemplate(write, "home.page.tmpl")
}

// About is the about page handler
func About(write http.ResponseWriter, request *http.Request){
	renderTemplate(write, "about.page.tmpl")
}