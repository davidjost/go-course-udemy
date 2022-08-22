package main

import (
	"errors"
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

func Divide(write http.ResponseWriter, request *http.Request)  {
	f, err := divideValues(100.0, 0.0)

	if err != nil {
		fmt.Fprintf(write, "cannot divide by 0")
		return
	}

	fmt.Fprintf(write, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

// lower case function names are like private funcs, only visible to the current package
func addValues(x, y int) int {
	return x + y 
}

func main() {
	
	// basic routing, for / call function Home
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application server on port %s", portNumber))
	
	_ = http.ListenAndServe(portNumber, nil)
}