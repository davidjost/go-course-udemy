package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		n, err := fmt.Fprintf(writer, "Hey web")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Number of Bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}