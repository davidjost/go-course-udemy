package main

import (
	"log"

	"github.com/davidjost/go-lesson-package/helpers"
)

const numPool = 1000


// func function_name( [parameter list] ) [return_types]
// the parameter this receives is bleh, of type channel, that channel only takes int as types.
func CalcValue(bleh chan int) {
	// write random number to randNum
	randNum := helpers.GenRandNum(numPool)
	// push randNum to the channel that comes from the parameter of the function
	bleh <- randNum
}

func main() {
	// channels lesson

	// open channel
	intChan := make(chan int)
	// closing a channel after use is required best practice
	defer close(intChan)

	// create extra go routine for concurrent execution
	go CalcValue(intChan)

	// num gets assigned whatever is coming out of the channel
	num := <-intChan
	log.Println("num is", num)
}
