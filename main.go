package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main() {
	// maps have to be set up with the shorthand syntax since a simple declaration without assignment results in a nil map. Nil maps cannot take any assignments.

	// map name := make(type[key type]value type)
	myMap := make(map[string]User)

	me := User {
		FirstName: "Trevor",
		LastName: "Philips",
	}

	myMap["me"] = me

	log.Println(myMap)
}
