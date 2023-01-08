package main

import (
	"fmt"
	"net/url"
)
const myurl string = "https://lco.dev:3000/learn?course=react&payment=123"

func main() {
	fmt.Println("Welcome to urls class")
	fmt.Println("url is :\n", myurl)

	//Parsing
	result, _:=url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams["course"][0]) // react

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",

	}
	fmt.Println("url :", partsOfUrl.String())
}


func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}
