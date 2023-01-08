package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to Slices class")
	
	var fruitList = []string{"Apple","Na","Ka"}
	fmt.Printf("Type is %T \n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println("List ", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("List ", fruitList)

	highScore := make([]int, 4)
	highScore[0] = 1
	highScore[1] = 2
	highScore[2] = 0
	highScore[3] = 9
	// highScore[4] = 888
	highScore = append(highScore, 44,45)
	fmt.Println("Scores", highScore)

	sort.Ints(highScore)
	fmt.Println("Scores", highScore)
	fmt.Println("Are  sorted", sort.IntsAreSorted(highScore))

	// remove value from slice based on index

	var nums = []int{1,2,3,5}
	fmt.Println("Numbers", nums)
	var index int = 2
	nums = append(nums[:index],nums[index+1:]...)
	fmt.Println("Numbers", nums)
}  