package main

import (
	"log"
)

func main() {
	// iterating over a slice of strings
	animals := []string{"horse", "cat", "dog", "fish", "cow"}

	// _ is a blank identifier, gets rid of the index variable
	for _, animal := range animals {
		log.Println(animal)
	}

	log.Println("-------------------")

	// maps are like objects with key/value pairs
	seaCreatures := make(map[string]string)

	seaCreatures["star"] = "patrick"
	seaCreatures["sponge"] = "bob"

	for animalType, name := range seaCreatures {
		log.Println(animalType, name)
	}

	log.Println("-------------------")
	
	// iterating over a string
	firstLine := "One upon a midnight dreary"
	
	// in Go, a string is a slice of bytes.
	for index, letter := range firstLine {
		log.Println(index, ":", letter)
	}

	log.Println("-------------------")

	// iterating over custom types
	type User struct {
		FirstName string
		LastName string
		Age int
	}

	// create variable as slice of type User
	var users []User
	// append into the slice the struct User and give it these values
	users = append(users, User{"John", "Milton", 45})
	users = append(users, User{"Will", "Wheaton", 22})
	users = append(users, User{"Mary", "Lamb", 24})
	users = append(users, User{"Mike", "Star", 66})

	for index, user := range users{
		log.Println(index, ":", user.FirstName, user.LastName, ",", user.Age)
	}
}