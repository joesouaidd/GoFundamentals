package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "joe",
			"last_name": "souaid",
			"hair_color": "brown",
			"has_dog": true
		}
	]`

	//write json to a stuct
	var Unmarshalled []Person

	// we transformed the string to byte in the function below
	error := json.Unmarshal([]byte(myJson), &Unmarshalled)
	if error != nil {
		fmt.Println("Error:", error)
	}
	fmt.Printf("unmarshalled: %v\n", Unmarshalled)

	//read json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = true

	mySlice = append(mySlice, m2)
	//marshall indent format the json nicely
	newJson, error := json.MarshalIndent(mySlice, "", "    ")
	if error != nil {
		fmt.Println("error", error)
	}

	fmt.Println(string(newJson))
}
