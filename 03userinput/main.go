package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "User input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza :")

	// comma ok // error ok

	input , _ := reader.ReadString('\n')
	fmt.Println(input)

	fmt.Println("Thanks for rating " ,input)
	fmt.Printf("Type rating %T" ,input)
}