package main

import (
	"fmt"
	"go-course-udemy/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	
	// basic routing, for / call function Home
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application server on port %s", portNumber))
	
	_ = http.ListenAndServe(portNumber, nil)
}