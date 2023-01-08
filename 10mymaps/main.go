package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to Maps class")
	
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of all langs", languages)

	fmt.Println("JS", languages["JS"])

	delete(languages,"JS")
	fmt.Println("JS", languages["JS"])
	fmt.Println("List of all langs", languages)

	// looping

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}  