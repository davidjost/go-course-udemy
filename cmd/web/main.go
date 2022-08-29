package main

import (
	"fmt"
	"go-course-udemy/pkg/config"
	"go-course-udemy/pkg/handlers"
	"go-course-udemy/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	// basic routing, for / call function Home
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application server on port %s", portNumber))
	
	_ = http.ListenAndServe(portNumber, nil)
}