package handlers

import (
	"go-course-udemy/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(write http.ResponseWriter, request *http.Request){
	render.RenderTemplate(write, "home.page.tmpl")
}

// About is the about page handler
func About(write http.ResponseWriter, request *http.Request){
	render.RenderTemplate(write, "about.page.tmpl")
}