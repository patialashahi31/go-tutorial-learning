package main

import "fmt"

// main is the first function will be called
func main() {
	fmt.Println("Hola")
	greeter()

	// Function inside func not allowed
	// func greeterTwo() {
		
	// }

	// result := adder(3,5)
	result,_ := proAdder(1,2,4)

	fmt.Println("Result", result)

}

func greeter() {
	fmt.Println("Hello from golang")
}

func adder(val1 int, val2 int) int{
	return val1 + val2
}

func proAdder(values ... int) (int , string){
	total := 0
	for _,value := range values {
		total += value
	}
	return total, "Hi"
}