package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davidjost/go-course-project-bookings/pkg/config"
	"github.com/davidjost/go-course-project-bookings/pkg/handlers"
	"github.com/davidjost/go-course-project-bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {
	// change this to true for prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session 

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