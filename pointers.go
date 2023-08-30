package main

import "log"

func main() {

	var myString string = "Green"

	log.Println("My screen is set to : ", myString)
	myString = changeUsingPointers(myString)
	log.Println("My screen is set to : ", myString)

	//OR

	//& is called ampersands
	changeUsingPointerss(&myString)
	log.Println("My screen is set to : ", myString)

}

//This approach is known as pass-by-value since a copy of the string is passed to the function.
//The original string in the main function is not modified by this function.

func changeUsingPointers(s string) string {

	newValue := "Red"
	s = newValue
	return s
}

// OR

//This approach is known as pass-by-reference since it modifies the original string through the pointer.
//The original string in the main function is modified by this function.
//pointer to a string
func changeUsingPointerss(s *string) {

	newValue := "Orange"

	//* is called asterisks
	*s = newValue

}
