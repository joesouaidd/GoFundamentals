package helpers

import (
	"math/rand"
	"time"
)

// Creating my own package
type SomeType struct {
	TypeName   string
	typeNumber int
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
