package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt stands for format
	fmt.Println("Hello World!")

	var whatToSay string 
	var i int

	whatToSay = "Goodbye cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to ", i)

	// fmt.Println(saySomething())
	// := var shorthand. Infers type from function type.
	whatWasSaid, thatOtherThing := saySomething()

	fmt.Println("The function returns:",whatWasSaid, thatOtherThing)

	// lesson 9: pointers
	// myString is set to Green and logged.
	myString := "Green"
	log.Println("myString is set to", myString)
	// change() is called with the position in memory where myString sits. The variable isn't getting a new value, what is at the position where the variable looks is being changed.
	// to reference a pointer, use &. To point to a location in memory, use *.
	changeUsingPointer(&myString)
	log.Println("after change through pointer", myString)
}

func saySomething() (string, string) {
	return "something", "else"
}

// lesson 9: pointers
func changeUsingPointer(s *string)  {
	log.Println("S is set to ", s)
	newValue := "Red"
	*s = newValue
}
