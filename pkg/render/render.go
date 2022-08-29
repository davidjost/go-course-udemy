package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "text/template"
)

func RenderTemplateTest(write http.ResponseWriter, tmpl string) {
	// we use _ here because template.ParseFiles has 2 returns: template and error, we only want the template
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(write, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

// package level variable for caching templates
var templateCache = make(map[string] *template.Template)

func RenderTemplate(write http.ResponseWriter, t string) {

	// this var is of the type that is *template.Template
	var tmpl *template.Template
	var err error

	// check if we already have the template in cache
	// we put the value of the parameter "t", which should be in the map "templateCache" into inMap
	// we ignore the first value since checking a map returns 2 things: key and value.
	_, inMap := templateCache[t]

	// if there is nothing in inMap
	if !inMap {
		// need to create the requested template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t) 

		 if err != nil{
			log.Println(err)
		 }
	} else {
		// we have the template
		log.Println("using cached template")
	}

	// put in tmpl the template by its name from the map templateCache
	tmpl = templateCache[t]

	if err != nil {
		log.Println(err)
	}

	err = tmpl.Execute(write, nil)
}

func createTemplateCache(templateName string) error {
	// create a slice of strings for the templates
	templates := []string{
		// Sprintf returns a string
		fmt.Sprintf("./templates/%s", templateName), "./templates/base.layout.tmpl",
	}	

	// parse the template. The ... spreads te slice contents out to individual strings.
	tmpl, err := template.ParseFiles(templates...)
	
	if err != nil {
		return err
	}

	// add template to cache (map)
	templateCache[templateName] = tmpl

	return nil
}