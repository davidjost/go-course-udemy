package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	
	// basic routing, for / call function Home
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application server on port %s", portNumber))
	
	_ = http.ListenAndServe(portNumber, nil)
}