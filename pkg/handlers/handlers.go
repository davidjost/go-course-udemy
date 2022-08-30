package handlers

import (
	"go-course-udemy/pkg/config"
	"go-course-udemy/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Respository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(write http.ResponseWriter, request *http.Request){
	render.RenderTemplate(write, "home.page.tmpl")
}

// About is the about page handler
func (m *Repository) About(write http.ResponseWriter, request *http.Request){
	render.RenderTemplate(write, "about.page.tmpl")
}