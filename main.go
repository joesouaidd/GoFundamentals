package main

import "fmt"

//main function is necessery always
func main() {

	fmt.Println("hello, world")

	var whatToSay string
	var i int

	whatToSay = "Goodbye"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to : ", i)
	whatWasSaid := saysomething()
	fmt.Println("The function returned ", whatWasSaid)

	_, whatWasSaid2 := saysomething2()
	fmt.Println("The second returned item from the second function is ", whatWasSaid2)
}

//return string
func saysomething() string {

	return "something"

}
func saysomething2() (string, string) {

	return "something", "else"

}
