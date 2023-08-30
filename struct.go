package main

import (
	"log"
	"time"
)

// if function is Declared starting a capital letter it is like defining it public in java
// by default it is private (when lower letter)
// same for variables
var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{FirstName: "Trevor", LastName: "Souaid",PhoneNumber: "+96170999666"}
	log.Println(user.FirstName, user.PhoneNumber)
}
