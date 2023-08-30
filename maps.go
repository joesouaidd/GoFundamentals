package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	//				  index   value
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])

	me := User{
		FirstName: "Joe",
		LastName:  "Souaid",
	}

	myMap2 := make(map[string]User)
	myMap2["Me"] = me

	log.Println(myMap2["Me"].FirstName)

	//maps are not sorted, must lookup by key

	//in maps we don't have to pass a pointer to map, we can just pass the map
	//it will remain constant everywhere, no matter where in program it is accessed

	// If I am not sure what type I will be storing in the MAP
	// I can use the type interface{}

	//myMap3 := make(map[string]interface{}) but then we need to cast it when we need to extract it

}
