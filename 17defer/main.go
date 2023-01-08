package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("World") // this line goes to last
	defer fmt.Println("World2") // LIFO
	fmt.Println("Welcome to defer class")
	 //Welcome to defer class
	// World
	myDefer()
}

func myDefer() {
	for i :=0; i<5 ; i++ {
		defer fmt.Println(i)
	}
}