package main

import "log"

func main() {

	var isTrue bool

	isTrue = false

	if isTrue {
		log.Println("is True is", isTrue)
	} else {
		log.Println("is True is", isTrue)
	}

	cat := "cat2"
	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 99

	if myNum == 99 && (cat == "cat2") {
		log.Println("&& done")
	}

	//switch statement

	myVar := "Cat"
	switch myVar {
	case "Cat":
		log.Println("cat")
	case "Dog":
		log.Println("Dog")

	default:
		log.Println("mostiquo")
	}

}
