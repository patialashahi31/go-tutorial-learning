package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to If else class")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular"
	} else if loginCount > 10 {
		result = "Watchout"
	} else {
		result =" 10 login count"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	}

	


}
