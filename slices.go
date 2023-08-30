package main

import (
	"log"
	"sort"
)

//slices are like arrays in other languages
//slices are appended by order

func main() {

	var mySlice []string

	mySlice = append(mySlice, "zz")
	mySlice = append(mySlice, "Hello")
	mySlice = append(mySlice, "World")

	log.Println(mySlice)
	sort.Strings(mySlice)
	log.Println(mySlice)

	numbers := [] int {1,2,3,4,5,6,7,8,9,10}

	log.Println(numbers)

	// 0 included, 2 excluded
	log.Println(numbers[0:2])

	

}
