package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files class")
	content := "This needs to go in file"
	file,err := os.Create("./content.txt")
	checkNilError(err)
	length, err := io.WriteString(file,content)
	checkNilError(err)
	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./content.txt")
}


func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data is : \n", string(databyte))
}

func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}