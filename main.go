package main

import (
	"fmt"
	"sort"
)

func main() {
	// slices are like arrays
	var mySlice []int
	
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	// sorting slices
	sort.Ints(mySlice)

	// log.Println(mySlice)

	// shorthand slices
	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	// slices like arrays start at 0, to limit the output use [:]
	// log.Println(numbers[4:10])

	// names := []string{"william", "jones", "pete"}

	// log.Println(names[1:2])

	// var numbers = make([]int,3,5)
  printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
