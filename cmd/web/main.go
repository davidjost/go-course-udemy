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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// calling this will give this program access to the app
	render.NewTemplates(&app)


	fmt.Println(fmt.Sprintf("Starting application server on port %s", portNumber))

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}