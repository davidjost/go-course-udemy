package handlers

import (
	"net/http"

	"github.com/davidjost/go-course-project-bookings/pkg/config"
	"github.com/davidjost/go-course-project-bookings/pkg/models"
	"github.com/davidjost/go-course-project-bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Respository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	// the & character is read as "address of", in this case it returns the address of Repository
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
	remoteIP := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(write, "home.page.tmpl", &models.TempateData{})	
}

// About is the about page handler
func (m *Repository) About(write http.ResponseWriter, request *http.Request){
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP


	// send the data to the template
	render.RenderTemplate(write, "about.page.tmpl", &models.TempateData{
		StringMap: stringMap,
	})
}