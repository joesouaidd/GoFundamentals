package main

import "log"

type myStruct struct {
	FirstName string
}

func (myStruct *myStruct) printFirstName(){
	log.Println(myStruct.FirstName)
} 

func main() {

	var myVar myStruct
	myVar.FirstName = "John"
	myVar2 := myStruct{FirstName: "Louis"}

	log.Println(myVar.FirstName)
	log.Println(myVar2.FirstName)
	myVar.printFirstName()
}