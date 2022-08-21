package helpers

import (
	"math/rand"
	"time"
)

func GenRandNum(number int) int {
	// we need to seed a new number to the rand function, we take the unix timecode for that
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(number)
	return value
}