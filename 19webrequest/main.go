package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web request class")
	response,err := http.Get(url)
	checkNilError(err)
	fmt.Println("content is: \n",response)
	defer response.Body.Close() // close 

	databytes,err:=ioutil.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println("string content:\n", string(databytes))
}


func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}