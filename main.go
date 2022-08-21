package main

import "log"

func main() {
	// slices are like arrays
	var mySlice []string
	
	mySlice = append(mySlice, "James")
	mySlice = append(mySlice, "Will")
	mySlice = append(mySlice, "Mary")

	log.Println(mySlice)
}
