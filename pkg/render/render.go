package render

import (
	"bytes"
	"go-course-udemy/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)


var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig)  {
	app = a
}

func RenderTemplate(write http.ResponseWriter, tmpl string) {

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

	err := t.Execute(buf, nil)
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

	// get all files named *.page.tmpl from ./templates/
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// range through all template files. Every entry in pages gets put into page.
	for _, page := range pages {
		// .Base strips off the path and leaves the filename
		fileName := filepath.Base(page)
		templateSet, err := template.New(fileName).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[fileName] = templateSet
	}

	return myCache, err
}