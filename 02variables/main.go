package main

import "fmt"

const LoginToken string = "sdgmnbsmgnzdgnmb"

func main()  {
	var username string = "u"
	fmt.Println(username)
	fmt.Printf("Type is : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type is : %T \n", isLoggedIn)


	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type is : %T \n", smallVal)


	var smallFloat float32 = 255.123465676768787979
	fmt.Println(smallFloat)

	var largeFloat float64 = 255.123465676768787979
	fmt.Println(largeFloat)

	// default values and aliases
	var defInt int
	fmt.Println(defInt)  // 0

	var defString string
	fmt.Println(defString)  // empty

	// implicit type

	var website = "http://ndnf.com" // string
	fmt.Println(website)

	// no var
	// most used style
	// only works inside a function
	noOfUsers := 3000
	fmt.Println(noOfUsers)

	fmt.Println(LoginToken)

}
