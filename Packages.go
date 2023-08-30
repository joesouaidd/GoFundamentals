package main

import (
	"fmt"

	helpers "github.com/joesouaidd/myniceprogram/Helpers"
)

func main() {
	fmt.Println("Hello")
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	fmt.Println(myVar.TypeName)
}
