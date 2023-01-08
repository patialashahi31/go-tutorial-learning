package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to Arrays class")
	
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "grapes"
	
	fmt.Println(" List", fruitList) //[apple banana  grapes]
	fmt.Println(" List", len(fruitList)) // 4

	var vegList =  [3]string{"potato","tomato","yo"}
	fmt.Println(" List", vegList)
}  