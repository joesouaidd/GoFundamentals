package main

import (
	"fmt"

	helpers "github.com/joesouaidd/myniceprogram/Helpers"
)

const max = 10
func CalculateValue (intChan chan int){

	randomNumber := helpers.RandomNumber(max)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)  //immidiatly close the channel when the function is done

	go CalculateValue(intChan)  //start as a go routine
	
	num := <-intChan
	fmt.Println(num)
}


