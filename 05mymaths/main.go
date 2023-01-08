package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main()  {
	fmt.Println("Welcome to maths class")

	// var  num1 int = 2
	// var num2 float64 = 4.5

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	randNum, _ := rand.Int(rand.Reader,big.NewInt(5))
	fmt.Println(randNum)



}  