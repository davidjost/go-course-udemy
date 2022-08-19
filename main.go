package main

import "log"

func main() {
	// maps have to be set up with the shorthand syntax since a simple declaration without assignment results in a nil map. Nil maps cannot take any assignments.

	// map name := make(type[key type]value type)
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	myMap["cat"] = "Clarance"

	log.Println(myMap)
}
