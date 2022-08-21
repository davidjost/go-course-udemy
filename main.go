package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// struct to manage json in go
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	// data source json from somewhere
	myJson := `
	[
		{
			"first_name" : "Clark",
			"last_name" : "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name" : "Bruce",
			"last_name" : "Wayne",
			"hair_color": "blonde",
			"has_dog": false
		}
	]`

	// create variable that is a slice of type Person, that will receive the json data
	var unmarshalled []Person

	// covert the json string with Unmarshal and write it to the place where the unmarshalled variable looks. JSON comes as a slice of bytes, thats why we need the []byte function.
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling the JSON", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write JSON from a struct
	// create var mySlice slice of type Person
	var mySlice []Person

	// create data for the slice
	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	// write data to slice
	mySlice = append(mySlice, m1)

	// create data for the slice
	var m2 Person
	m2.FirstName = "John"
	m2.LastName = "Wayne"
	m2.HairColor = "brown"
	m2.HasDog = false

	// write data to slice
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("Error marshalling the JSON", err)
	}

	fmt.Println(string(newJson))
}
