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
	NumberOfTeeth int
}

func main() {
	dog := Dog{"tequila", "Malinous"}
	PrintInfo(&dog)

	gorilla := Gorilla{"bigbiss", "black", 38}
	PrintInfo(&gorilla)

}

//in this function we've put the input as an interface to force having the object certain functions
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), " and has", a.NumberOfLegs(), " legs")
}

func (d *Dog) Says() string {
	return "Woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "wooo"
}

func (g *Gorilla) NumberOfLegs() int {
	return 4
}
