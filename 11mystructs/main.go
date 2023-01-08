package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Struct class")

	// no inheritance
	// no super or parent

	myUser := User{"Tej","tej@a.com",true,26}
	fmt.Println("User", myUser) //{Tej tej@a.com true 26}
	fmt.Println("Age", myUser.Age)
	fmt.Printf("More details is %+v \n",myUser) // {Name:Tej Email:tej@a.com Status:true Age:26} 

}

// U means public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
