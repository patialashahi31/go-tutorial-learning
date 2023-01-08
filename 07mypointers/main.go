package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to Pointers class")

	// var ptr *int
	// fmt.Println("Value of ptr is" , ptr) //<nil>

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of ptr is ," , ptr)
	fmt.Println("Value is ," , *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is ," , myNumber) //46
}  