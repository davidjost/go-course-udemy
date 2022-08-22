package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(write http.ResponseWriter, request *http.Request){
	fmt.Fprintf(write, "This is the home page")
}

// About is the about page handler
func About(write http.ResponseWriter, request *http.Request){
	sum := addValues(2, 2)
	// the underscores are used to ignore expected values to get the program to run
	_, _ = fmt.Fprintf(write, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// lower case function names are like private funcs, only visible to the current package
func addValues(x, y int) int {
	return x + y 
}

func main() {
	
	// basic routing, for / call function Home
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application server on port %s", portNumber))
	
	_ = http.ListenAndServe(portNumber, nil)
}