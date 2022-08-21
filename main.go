package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberofTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "Collie",
	}

	PrintInfo(&dog)
	
	gorilla := Gorilla{
		Name: "Bobo",
		Color: "Gray",
		NumberofTeeth: 24,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"	
}

func (d *Dog) NumberOfLegs() int {
	return 4	
}

func (d *Gorilla) Says() string {
	return "Ugh"	
}

func (d *Gorilla) NumberOfLegs() int {
	return 2	
}