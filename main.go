package main

import (
	"log"
)

func main() {
	myNum := 100
	buul := false

	// logical operators
	if myNum > 99 && !buul {
		log.Println("myNum is above 99 and buul is", buul)
	}

	// switch

	myVar := "fish"

	switch myVar {
	case "dog":
		log.Println("It's a dog.")
	case "cat":
		log.Println("meow")
	default:
		log.Println("default")
	}

}