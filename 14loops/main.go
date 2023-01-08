package main

import "fmt"
	

func main() {
	fmt.Println("Welcome to Loop class")
	
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println("Days ", days)

	// for i:=0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range(days) {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println(index,day)
	// }

	// while loop
	rougueValue := 1
	
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println(rougueValue)
		rougueValue ++
	}

	lco:
	fmt.Println("Jump at ")

}
