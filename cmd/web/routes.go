package main

import (
	"net/http"

	"github.com/davidjost/go-course-project-bookings/pkg/config"
	"github.com/davidjost/go-course-project-bookings/pkg/handlers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/avail", handlers.Repo.Avail)
	mux.Get("/availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/make-reservation", handlers.Repo.Reservation)

	// fileserver
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	
	return mux
}
