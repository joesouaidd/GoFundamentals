package main

import (
	"fmt"
	"log"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	animals := []string{"dog", "fish", "cat", "horse", "cow"}

	// range returns, first one the index, the second one is the value at that index
	// the _ means ignore this value
	for _, animal := range animals {
		log.Println(animal)

	}

	//Loop a map
	foods := make(map[string]int)
	foods["twixx"] = 15
	foods["oreo"] = 20

	for _, food := range foods {
		fmt.Print(food, " ")
	}

	//loop a string
	var line = "Once upon a time"
	for i, l := range line {
		fmt.Println(i, ";", l)
	}

	//loop a struct
	type User struct {
		FirstName string
		LastName  string
		Age       int
	}

	var users []User
	users = append(users, User{"Joe", "Souaid", 99})
	users = append(users, User{"Alex", "zi", 20})

	for _, l := range users {
		fmt.Println(l.FirstName, l.LastName, l.Age)
	}
}
