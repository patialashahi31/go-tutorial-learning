package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


func main() {
	fmt.Println("Welcome to web verb class")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
	
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err:= http.Get(myurl)
	checkNilError(err)
	defer response.Body.Close()

	fmt.Println("Status code", response.StatusCode)
	fmt.Println("Content length", response.ContentLength)
	var responseString strings.Builder
	content, _:=ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	responseString.Write(content)
	fmt.Println(responseString.String())


}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
	{
		"course": "Go lang",
		"price": 0
	}
	`)

	response, _:=http.Post(myurl,"application/json",requestBody)
	
	defer response.Body.Close()
	content, _:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname", "tej")

	response, _:=http.PostForm(myurl,data)
	
	defer response.Body.Close()
	content, _:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}


func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}
