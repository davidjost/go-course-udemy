package main

import "log"

/*
https://www.tutorialspoint.com/go/go_structures.htm
Structures are used to represent a record. Suppose you want to keep track of the books in a library. You might want to track the following attributes of each book −

		Title
		Author
		Subject
		Book ID

In such a scenario, structures are highly useful.
*/
type myStruct struct {
	FirstName string
}

// assigning functions to structs, called a receiver, it ties the function to the struct myStruct
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	// single struct value assignment
	var myVar myStruct
	myVar.FirstName = "John"

	// shorthand struct value assignment
	myVar2 := myStruct{
		FirstName: "Mary",
	}

	// myVar is set to the myStruct type, which also has the function attached to it. It's possible to write logic for the data in the struct attached to the same struct, which groups source and manipulation nicely.
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
