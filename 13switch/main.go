package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch class")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(" Value of dice is ," , diceNumber)
	
	switch diceNumber {
	case 1:
		fmt.Println(" Dice is 1 and you can open")
	case 2:
		fmt.Println(" You can move to 2")
	case 3:
		fmt.Println(" You can move to 3")
		fallthrough
	case 4:
		fmt.Println(" You can move to 4")
	case 5:
		fmt.Println(" You can move to 5")
	case 6:
		fmt.Println(" You can move to 6")
	default:
		fmt.Println(" What was that")
	}

}
