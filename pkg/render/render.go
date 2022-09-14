package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/davidjost/go-course-project-bookings/pkg/config"
	"github.com/davidjost/go-course-project-bookings/pkg/models"
)


var app *config.AppConfig

// NewTemplates func sets the config for the template package
func NewTemplates(a *config.AppConfig)  {
	app = a
}

func AddDefaultData(td *models.TempateData) *models.TempateData {
	return td
}

func RenderTemplate(write http.ResponseWriter, tmpl string, td *models.TempateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get the requested template from cache. 2 variable declarations for the return values of createTemplateCache(), which is a map: key and value pair.
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(write)
	if err != nil {
		log.Println(err)
	}
}

// uppercase func names! Otherwise, the main.go cannot see the function in this render package.
func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache is a map of strings of type *template.Template and also is empty due to the {}
	myCache := map[string]*template.Template{}

	// get all files named *.page.html from ./templates/
	pages, err := filepath.Glob("./templates/*.page.html")
	// log.Println(pages)

	if err != nil {
		// log.Println("Error 1")
		return myCache, err
	}

	// range through all template files. Every entry in pages gets put into page.
	for _, page := range pages {
		// .Base strips off the path and leaves the filename
		fileName := filepath.Base(page)
		// log.Println("fileName is:", fileName)
		// log.Println("page is:", page)
		templateSet, err := template.New(fileName).ParseFiles(page)
		// log.Println("templateSet is:", templateSet)

		if err != nil {
			// log.Println("Error 2")
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		
		if err != nil {
			// log.Println("Error 3")
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				// log.Println("Error 4")
				return myCache, err
			}
		}

		myCache[fileName] = templateSet
	}

	return myCache, err
}