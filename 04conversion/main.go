package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "User input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza between 1 and 5:")

	// comma ok // error ok

	input , _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating " ,input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 to rating", numRating + 1)
	}
}